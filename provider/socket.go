package provider

import (
	"crypto/tls"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/andig/evcc/util"
	"github.com/andig/evcc/util/jq"
	"github.com/gorilla/websocket"
	"github.com/itchyny/gojq"
)

// Socket implements websocket request provider
type Socket struct {
	*util.HTTPHelper
	mux     *util.Waiter
	url     string
	headers map[string]string
	scale   float64
	jq      *gojq.Query
	val     interface{}
}

// NewSocketProviderFromConfig creates a HTTP provider
func NewSocketProviderFromConfig(log *util.Logger, other map[string]interface{}) *Socket {
	cc := struct {
		URI      string
		Headers  map[string]string
		Jq       string
		Scale    float64
		Insecure bool
		Auth     Auth
		Timeout  time.Duration
	}{}
	util.DecodeOther(log, other, &cc)

	logger := util.NewLogger("ws")

	p := &Socket{
		HTTPHelper: util.NewHTTPHelper(logger),
		mux:        util.NewWaiter(cc.Timeout, func() { logger.TRACE.Println("wait for initial value") }),
		url:        cc.URI,
		headers:    cc.Headers,
		scale:      cc.Scale,
	}

	// handle basic auth
	if cc.Auth.Type != "" {
		if p.headers == nil {
			p.headers = make(map[string]string)
		}
		NewAuth(log, cc.Auth, p.headers)
	}

	// ignore the self signed certificate
	if cc.Insecure {
		customTransport := http.DefaultTransport.(*http.Transport).Clone()
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		p.HTTPHelper.Client.Transport = customTransport
	}

	if cc.Jq != "" {
		op, err := gojq.Parse(cc.Jq)
		if err != nil {
			log.FATAL.Fatalf("config: invalid jq query: %s", p.jq)
		}

		p.jq = op
	}

	go p.listen()

	return p
}

func (p *Socket) listen() {
	log := p.HTTPHelper.Log

	headers := make(http.Header)
	for k, v := range p.headers {
		headers.Set(k, v)
	}

	for {
		client, _, err := websocket.DefaultDialer.Dial(p.url, headers)
		if err != nil {
			log.ERROR.Println("dial:", err)
		}

		for {
			_, b, err := client.ReadMessage()
			if err != nil {
				log.TRACE.Println("read:", err)
				_ = client.Close()
				break
			}

			log.TRACE.Printf("recv: %s", b)

			p.mux.Lock()
			if p.jq != nil {
				v, err := jq.Query(p.jq, b)
				if err == nil {
					p.val = v
					p.mux.Update()
				} else {
					log.WARN.Printf("invalid: %s", string(b))
				}
			} else {
				p.val = string(b)
				p.mux.Update()
			}
			p.mux.Unlock()
		}
	}
}

func (p *Socket) hasValue() (interface{}, error) {
	elapsed := p.mux.LockWithTimeout()
	defer p.mux.Unlock()

	if elapsed > 0 {
		return nil, fmt.Errorf("outdated: %v", elapsed.Truncate(time.Second))
	}

	return p.val, nil
}

// StringGetter sends string request
func (p *Socket) StringGetter() (string, error) {
	v, err := p.hasValue()
	if err != nil {
		return "", err
	}

	return jq.String(v)
}

// FloatGetter parses float from string getter
func (p *Socket) FloatGetter() (float64, error) {
	v, err := p.hasValue()
	if err != nil {
		return 0, err
	}

	// v is always string when jq not used
	if p.jq == nil {
		v, err = strconv.ParseFloat(v.(string), 64)
		if err != nil {
			return 0, err
		}
	}

	f, err := jq.Float64(v)
	return f * p.scale, err
}

// IntGetter parses int64 from float getter
func (p *Socket) IntGetter() (int64, error) {
	v, err := p.hasValue()
	if err != nil {
		return 0, err
	}

	// v is always string when jq not used
	if p.jq == nil {
		v, err = strconv.ParseInt(v.(string), 10, 64)
		if err != nil {
			return 0, err
		}
	}

	i, err := jq.Int64(v)
	f := float64(i) * p.scale

	return int64(math.Round(f)), err
}

// BoolGetter parses bool from string getter
func (p *Socket) BoolGetter() (bool, error) {
	v, err := p.hasValue()
	if err != nil {
		return false, err
	}

	// v is always string when jq not used
	if p.jq == nil {
		v = util.Truish(v.(string))
	}

	return jq.Bool(v)
}
