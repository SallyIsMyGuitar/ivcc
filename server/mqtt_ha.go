package server

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/evcc-io/evcc/core/site"
)

type HADeviceDef struct {
	Identifiers  string `json:"identifiers"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}

type HAEntityDef struct {
	Device   HADeviceDef `json:"device"`
	Name     string      `json:"name"`
	UniqueId string      `json:"unique_id"`

	StateTopic        string `json:"state_topic,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`

	Options []string `json:"options,omitempty"`

	Min float64 `json:"min"`
	Max float64 `json:"max"`

	StateClass string `json:"state_class,omitempty"`
}

func (m *MQTT) haPublishBaseDeviceDef(site site.API, loadPoint *int, valueName string) HAEntityDef {
	siteId := strings.ReplaceAll(strings.ToLower(site.GetTitle()), " ", "_") // TODO find better site ID
	uid := "evcc_" + siteId + "_"
	if loadPoint != nil {
		uid += fmt.Sprintf("lp%d_", *loadPoint+1)
	}
	uid += valueName

	var name string
	if loadPoint == nil {
		name = fmt.Sprintf("EVCC %s %s", site.GetTitle(), valueName)
	} else {
		name = fmt.Sprintf("EVCC %s Loadpoint %d %s", site.GetTitle(), *loadPoint+1, valueName)
	}

	return HAEntityDef{
		Device: HADeviceDef{
			Identifiers:  fmt.Sprintf("evcc_%s", siteId),
			Manufacturer: "EVCC",
			Name:         site.GetTitle(),
		},
		Name:              name,
		UniqueId:          uid,
		AvailabilityTopic: m.root + "/status",
	}
}

func (m *MQTT) haSendEntityDef(topic string, def HAEntityDef) {
	jsonData, _ := json.MarshalIndent(def, "", "  ")
	token := m.Handler.Client.Publish(topic, m.Handler.Qos, true, jsonData)
	go m.Handler.WaitForToken(token)

	// mark as know to sensors publish only once
	m.haKnownSensors[def.StateTopic] = struct{}{}
}

func (m *MQTT) haPublishDiscoverSensors(site site.API, loadPoint *int, valueName, stateTopic string, val interface{}) {
	if _, ok := m.haKnownSensors[stateTopic]; ok {
		return
	}

	entityDef := m.haPublishBaseDeviceDef(site, loadPoint, valueName)
	entityDef.StateTopic = stateTopic

	switch val.(type) {
	case float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, []float64:
		entityDef.StateClass = "measurement"
	}

	m.haSendEntityDef("homeassistant/sensor/"+entityDef.UniqueId+"/config", entityDef)
}

func (m *MQTT) haPublishDiscoverSelect(site site.API, loadPoint *int, valueName, stateTopic string, options []string) {
	entityDef := m.haPublishBaseDeviceDef(site, loadPoint, valueName)

	entityDef.StateTopic = stateTopic
	entityDef.CommandTopic = stateTopic + "/set"
	entityDef.Options = options

	m.haSendEntityDef("homeassistant/select/"+entityDef.UniqueId+"/config", entityDef)
}

func (m *MQTT) haPublishDiscoverNumber(site site.API, loadPoint *int, valueName, stateTopic string, min, max float64) {
	entityDef := m.haPublishBaseDeviceDef(site, loadPoint, valueName)

	entityDef.StateTopic = stateTopic
	entityDef.CommandTopic = stateTopic + "/set"
	entityDef.Min = min
	entityDef.Max = max

	m.haSendEntityDef("homeassistant/number/"+entityDef.UniqueId+"/config", entityDef)
}
