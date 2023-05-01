package server

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/evcc-io/evcc/util"
	"github.com/kr/pretty"
	"nhooyr.io/websocket"
)

const (
	// Time allowed to write a message to the peer
	socketWriteTimeout = 10 * time.Second
)

// SocketClient is a middleman between the websocket connection and the hub.
type SocketClient struct {
	// Buffered channel of outbound messages.
	send chan []byte
}

// ServeWebsocket handles websocket requests from the peer.
func ServeWebsocket(hub *SocketHub, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.ERROR.Println(err)
		return
	}

	client := &SocketClient{send: make(chan []byte, 256)}
	hub.register <- client

	defer func() {
		conn.Close(websocket.StatusNormalClosure, "done")
		hub.unregister <- client
	}()

	for msg := range client.send {
		// ctx, cancel := context.WithTimeout(context.Background(), socketWriteTimeout)
		err := conn.Write(context.Background(), websocket.MessageText, msg)
		// cancel()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// SocketHub maintains the set of active clients and broadcasts messages to the
// clients.
type SocketHub struct {
	mu sync.RWMutex

	// Registered clients.
	clients map[*SocketClient]bool

	// Register requests from the clients.
	register chan *SocketClient

	// Unregister requests from clients.
	unregister chan *SocketClient
}

// NewSocketHub creates a web socket hub that distributes meter status and
// query results for the ui or other clients
func NewSocketHub() *SocketHub {
	return &SocketHub{
		register:   make(chan *SocketClient),
		unregister: make(chan *SocketClient),
		clients:    make(map[*SocketClient]bool),
	}
}

func encode(v interface{}) (string, error) {
	var s string
	switch val := v.(type) {
	case time.Time:
		if val.IsZero() {
			s = "null"
		} else {
			s = fmt.Sprintf(`"%s"`, val.Format(time.RFC3339))
		}
	case time.Duration:
		// must be before stringer to convert to seconds instead of string
		s = fmt.Sprintf("%d", int64(val.Seconds()))
	case float64:
		if math.IsNaN(val) {
			s = "null"
		} else {
			s = fmt.Sprintf("%.5g", val)
		}
	default:
		if b, err := json.Marshal(v); err == nil {
			s = string(b)
		} else {
			return "", err
		}
	}
	return s, nil
}

func kv(p util.Param) string {
	val, err := encode(p.Val)
	if err != nil {
		panic(err)
	}

	if p.Key == "" && val == "" {
		log.ERROR.Printf("invalid key/val for %+v %# v, please report to https://github.com/evcc-io/evcc/issues/6439", p, pretty.Formatter(p.Val))
		return "\"foo\":\"bar\""
	}

	var msg strings.Builder
	msg.WriteString("\"")
	if p.Loadpoint != nil {
		msg.WriteString(fmt.Sprintf("loadpoints.%d.", *p.Loadpoint))
	}
	msg.WriteString(p.Key)
	msg.WriteString("\":")
	msg.WriteString(val)

	return msg.String()
}

func (h *SocketHub) welcome(client *SocketClient, params []util.Param) {
	h.mu.Lock()
	h.clients[client] = true
	h.mu.Unlock()

	var msg strings.Builder
	msg.WriteString("{")
	for _, p := range params {
		if msg.Len() > 1 {
			msg.WriteString(",")
		}
		msg.WriteString(kv(p))
	}
	msg.WriteString("}")

	select {
	case client.send <- []byte(msg.String()):
	default:
		close(client.send)
	}
}

func (h *SocketHub) broadcast(p util.Param) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.clients) > 0 {
		msg := "{" + kv(p) + "}"

		for client := range h.clients {
			select {
			case client.send <- []byte(msg):
			default:
				h.unregister <- client
			}
		}
	}
}

// Run starts data and status distribution
func (h *SocketHub) Run(in <-chan util.Param, cache *util.Cache) {
	for {
		select {
		case client := <-h.register:
			h.welcome(client, cache.All())
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				close(client.send)
				delete(h.clients, client)
			}
			h.mu.Unlock()
		case msg, ok := <-in:
			if !ok {
				return // break if channel closed
			}
			h.broadcast(msg)
		}
	}
}
