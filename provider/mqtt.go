package provider

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/andig/evcc/api"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	publishTimeout = 2 * time.Second
	waitTimeout    = 50 * time.Millisecond // polling interval when waiting for initial value
)

// MqttClient is a paho publisher
type MqttClient struct {
	log      *api.Logger
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
	log := api.NewLogger("mqtt")
	log.INFO.Printf("connecting %s at %s", clientID, broker)

	mc := &MqttClient{
		log:      log,
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
	m.log.WARN.Printf("%s connection lost: %v", m.broker, reason.Error())
}

// ConnectionHandler restores listeners
func (m *MqttClient) ConnectionHandler(client mqtt.Client) {
	m.log.TRACE.Printf("%s connected", m.broker)

	m.mux.Lock()
	defer m.mux.Unlock()

	for topic, l := range m.listener {
		m.log.TRACE.Printf("%s subscribe %s", m.broker, topic)
		go m.listen(topic, l)
	}
}

// Listen validates uniqueness and registers and attaches listener
func (m *MqttClient) Listen(topic string, callback func(string)) {
	m.mux.Lock()
	if _, ok := m.listener[topic]; ok {
		m.log.FATAL.Fatalf("%s: duplicate listener not allowed", topic)
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
		log:        m.log,
		topic:      topic,
		multiplier: multiplier,
		timeout:    timeout,
	}

	m.Listen(topic, h.Receive)
	return h.floatGetter
}

// IntGetter creates handler for int64 from MQTT topic that returns cached value
func (m *MqttClient) IntGetter(topic string, multiplier int64, timeout time.Duration) IntGetter {
	h := &msgHandler{
		log:        m.log,
		topic:      topic,
		multiplier: float64(multiplier),
		timeout:    timeout,
	}

	m.Listen(topic, h.Receive)
	return h.intGetter
}

// StringGetter creates handler for string from MQTT topic that returns cached value
func (m *MqttClient) StringGetter(topic string, timeout time.Duration) StringGetter {
	h := &msgHandler{
		log:     m.log,
		topic:   topic,
		timeout: timeout,
	}

	m.Listen(topic, h.Receive)
	return h.stringGetter
}

// BoolGetter creates handler for string from MQTT topic that returns cached value
func (m *MqttClient) BoolGetter(topic string, timeout time.Duration) BoolGetter {
	h := &msgHandler{
		log:     m.log,
		topic:   topic,
		timeout: timeout,
	}

	m.Listen(topic, h.Receive)
	return h.boolGetter
}

// formatValue formats a message template of returns the value formatted as %v is template is empty
func (m *MqttClient) formatValue(param, message string, v interface{}) (string, error) {
	if message == "" {
		return fmt.Sprintf("%v", v), nil
	}

	return replaceFormatted(message, map[string]interface{}{
		param: v,
	})
}

// IntSetter publishes topic with parameter replaced by int value
func (m *MqttClient) IntSetter(param, topic, message string) IntSetter {
	return func(v int64) error {
		payload, err := m.formatValue(param, message, v)
		if err != nil {
			return err
		}

		m.log.TRACE.Printf("send %s: '%s'", topic, payload)
		token := m.Client.Publish(topic, m.qos, false, payload)
		if token.WaitTimeout(publishTimeout) {
			return token.Error()
		}

		return fmt.Errorf("%s send timeout", topic)
	}
}

// BoolSetter invokes script with parameter replaced by bool value
func (m *MqttClient) BoolSetter(param, topic, message string) BoolSetter {
	return func(v bool) error {
		payload, err := m.formatValue(param, message, v)
		if err != nil {
			return err
		}

		m.log.TRACE.Printf("send %s: '%s'", topic, payload)
		token := m.Client.Publish(topic, m.qos, false, payload)
		if token.WaitTimeout(publishTimeout) {
			return token.Error()
		}

		return fmt.Errorf("%s send timeout", topic)
	}
}

// WaitForToken synchronously waits until token operation completed
func (m *MqttClient) WaitForToken(token mqtt.Token) {
	if token.WaitTimeout(publishTimeout) {
		if token.Error() != nil {
			m.log.ERROR.Printf("error: %s", token.Error())
		}
	} else {
		m.log.DEBUG.Println("timeout")
	}
}

type msgHandler struct {
	log        *api.Logger
	once       sync.Once
	mux        sync.Mutex
	updated    time.Time
	timeout    time.Duration
	multiplier float64
	topic      string
	payload    string
}

func (h *msgHandler) Receive(payload string) {
	h.log.TRACE.Printf("recv %s: '%s'", h.topic, payload)

	h.mux.Lock()
	defer h.mux.Unlock()

	h.payload = payload
	h.updated = time.Now()
}

func (h *msgHandler) waitForInitialValue() {
	h.mux.Lock()
	defer h.mux.Unlock()

	if h.updated.IsZero() {
		h.log.TRACE.Printf("%s wait for initial value", h.topic)

		// wait for initial update
		for h.updated.IsZero() {
			h.mux.Unlock()
			time.Sleep(waitTimeout)
			h.mux.Lock()
		}
	}
}

func (h *msgHandler) floatGetter() (float64, error) {
	h.once.Do(h.waitForInitialValue)
	h.mux.Lock()
	defer h.mux.Unlock()

	if elapsed := time.Since(h.updated); h.timeout != 0 && elapsed > h.timeout {
		return 0, fmt.Errorf("%s outdated: %v", h.topic, elapsed.Truncate(time.Second))
	}

	val, err := strconv.ParseFloat(h.payload, 64)
	if err != nil {
		return 0, fmt.Errorf("%s invalid: '%s'", h.topic, h.payload)
	}

	return h.multiplier * val, nil
}

func (h *msgHandler) intGetter() (int64, error) {
	f, err := h.floatGetter()
	return int64(math.Round(f)), err
}

func (h *msgHandler) stringGetter() (string, error) {
	h.once.Do(h.waitForInitialValue)
	h.mux.Lock()
	defer h.mux.Unlock()

	if elapsed := time.Since(h.updated); h.timeout != 0 && elapsed > h.timeout {
		return "", fmt.Errorf("%s outdated: %v", h.topic, elapsed.Truncate(time.Second))
	}

	return string(h.payload), nil
}

func (h *msgHandler) boolGetter() (bool, error) {
	h.once.Do(h.waitForInitialValue)
	h.mux.Lock()
	defer h.mux.Unlock()

	if elapsed := time.Since(h.updated); h.timeout != 0 && elapsed > h.timeout {
		return false, fmt.Errorf("%s outdated: %v", h.topic, elapsed.Truncate(time.Second))
	}

	return truish(string(h.payload)), nil
}
