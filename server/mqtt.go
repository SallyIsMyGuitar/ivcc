package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/core"
	"github.com/andig/evcc/provider"
	"github.com/andig/evcc/util"
)

// MQTT is the MQTT server. It uses the MQTT client for publishing.
type MQTT struct {
	Handler *provider.MqttClient
}

func (m *MQTT) encode(v interface{}) string {
	var s string
	switch val := v.(type) {
	case time.Time:
		s = strconv.FormatInt(val.Unix(), 10)
	case time.Duration:
		// must be before stringer to convert to seconds instead of string
		s = fmt.Sprintf("%d", int64(val.Seconds()))
	case fmt.Stringer, string:
		s = fmt.Sprintf("%s", val)
	case float64:
		s = fmt.Sprintf("%.3f", val)
	default:
		s = fmt.Sprintf("%v", val)
	}
	return s
}

func (m *MQTT) publish(topic string, retained bool, payload interface{}) {
	token := m.Handler.Client.Publish(topic, m.Handler.Qos, retained, m.encode(payload))
	go m.Handler.WaitForToken(token)
}

type apiHandler interface {
	SetMode(api.ChargeMode)
	SetTargetSoC(int)
}

func (m *MQTT) listenSetters(topic string, apiHandler apiHandler) {
	m.Handler.Listen(topic+"/mode/set", func(payload string) {
		apiHandler.SetMode(api.ChargeMode(payload))
	})
	m.Handler.Listen(topic+"/targetsoc/set", func(payload string) {
		soc, err := strconv.Atoi(payload)
		if err == nil {
			apiHandler.SetTargetSoC(soc)
		}
	})
}

// Run starts the MQTT publisher for the MQTT API
func (m *MQTT) Run(root string, site *core.Site, in <-chan util.Param) {
	topic := fmt.Sprintf("%s/site", root)
	m.listenSetters(topic, site)

	// number of loadpoints
	topic = fmt.Sprintf("%s/loadpoints", root)
	m.publish(topic, true, len(site.LoadPoints()))

	for id, lp := range site.LoadPoints() {
		topic := fmt.Sprintf("%s/loadpoints/%d", root, id)
		m.listenSetters(topic, lp)
	}

	// alive indicator
	updated := time.Now().Unix()
	m.publish(fmt.Sprintf("%s/updated", root), true, updated)

	for p := range in {
		topic = fmt.Sprintf("%s/site", root)
		if p.LoadPoint != nil {
			topic = fmt.Sprintf("%s/loadpoints/%d", root, *p.LoadPoint)
		}

		// alive indicator
		if now := time.Now().Unix(); now != updated {
			updated = now
			m.publish(fmt.Sprintf("%s/updated", root), true, updated)
		}

		// value
		topic += "/" + strings.ToLower(p.Key)
		m.publish(topic, false, p.Val)
	}
}
