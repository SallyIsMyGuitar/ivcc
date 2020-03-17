package provider

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/andig/evcc/api"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const publishTimeout = 2 * time.Second

var mlog = api.NewLogger("mqtt")

// MqttClient is a paho publisher
type MqttClient struct {
	mux      sync.Mutex
	Client   mqtt.Client
	broker   string
	qos      byte
	listener map[string]func(string)
}

// NewMqttClient creates new publisher for paho
func NewMqttClient(
	broker string,
	user string,
	password string,
	clientID string,
	qos byte,
) *MqttClient {
	mlog.INFO.Printf("connecting %s at %s", clientID, broker)

	mc := &MqttClient{
		broker:   broker,
		qos:      qos,
		listener: make(map[string]func(string)),
	}

	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetUsername(user)
	options.SetPassword(password)
	options.SetClientID(clientID)
	options.SetCleanSession(true)
	options.SetAutoReconnect(true)
	options.SetOnConnectHandler(mc.ConnectionHandler)
	options.SetConnectionLostHandler(mc.ConnectionLostHandler)

	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.FATAL.Fatalf("error connecting: %s", token.Error())
	}

	mc.Client = client
	return mc
}

// ConnectionLostHandler logs cause of connection loss as warning
func (m *MqttClient) ConnectionLostHandler(client mqtt.Client, reason error) {
	mlog.WARN.Printf("%s connection lost: %v", m.broker, reason.Error())
}

// ConnectionHandler restores listeners
func (m *MqttClient) ConnectionHandler(client mqtt.Client) {
	mlog.TRACE.Printf("%s connected", m.broker)

	m.mux.Lock()
	defer m.mux.Unlock()

	for topic, l := range m.listener {
		mlog.TRACE.Printf("%s subscribe %s", m.broker, topic)
		go m.listen(topic, l)
	}
}

// Listen validates uniqueness and registers and attaches listener
func (m *MqttClient) Listen(topic string, callback func(string)) {
	m.mux.Lock()
	if _, ok := m.listener[topic]; ok {
		mlog.FATAL.Fatalf("%s: duplicate listener not allowed", topic)
	}
	m.listener[topic] = callback
	m.mux.Unlock()

	m.listen(topic, callback)
}

// listen attaches listener to topic
func (m *MqttClient) listen(topic string, callback func(string)) {
	token := m.Client.Subscribe(topic, m.qos, func(c mqtt.Client, msg mqtt.Message) {
		s := string(msg.Payload())
		if len(s) > 0 {
			callback(s)
		}
	})
	m.WaitForToken(token)
}

// FloatGetter creates handler for float64 from MQTT topic that returns cached value
func (m *MqttClient) FloatGetter(topic string, multiplier float64, timeout time.Duration) FloatGetter {
	h := &msgHandler{
		topic:      topic,
		multiplier: multiplier,
		timeout:    timeout,
	}

	m.Listen(topic, h.Receive)
	return h.floatGetter
}

// IntGetter creates handler for int64 from MQTT topic that returns cached value
func (m *MqttClient) IntGetter(topic string, timeout time.Duration) IntGetter {
	h := &msgHandler{
		topic:   topic,
		timeout: timeout,
	}

	m.Listen(topic, h.Receive)
	return h.intGetter
}

// WaitForToken synchronously waits until token operation completed
func (m *MqttClient) WaitForToken(token mqtt.Token) {
	if token.WaitTimeout(publishTimeout) {
		if token.Error() != nil {
			mlog.ERROR.Printf("error: %s", token.Error())
		}
	} else {
		mlog.DEBUG.Println("timeout")
	}
}

type msgHandler struct {
	mux        sync.Mutex
	updated    time.Time
	timeout    time.Duration
	multiplier float64
	topic      string
	payload    string
}

func (h *msgHandler) Receive(payload string) {
	mlog.TRACE.Printf("recv %s: '%s'", h.topic, payload)

	h.mux.Lock()
	defer h.mux.Unlock()

	h.payload = payload
	h.updated = time.Now()
}

func (h *msgHandler) floatGetter() (float64, error) {
	h.mux.Lock()
	defer h.mux.Unlock()

	if time.Since(h.updated) > h.timeout {
		return 0, fmt.Errorf("%s outdated: %v", h.topic, time.Since(h.updated))
	}

	val, err := strconv.ParseFloat(h.payload, 64)
	if err != nil {
		return 0, fmt.Errorf("%s invalid: '%s'", h.topic, h.payload)
	}

	return h.multiplier * val, nil
}

func (h *msgHandler) intGetter() (int64, error) {
	h.mux.Lock()
	defer h.mux.Unlock()

	if time.Since(h.updated) > h.timeout {
		return 0, fmt.Errorf("%s outdated: %v", h.topic, time.Since(h.updated))
	}

	val, err := strconv.ParseInt(h.payload, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s invalid: '%s'", h.topic, h.payload)
	}

	return val, nil
}
