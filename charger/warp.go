package charger

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/provider"
	"github.com/andig/evcc/provider/mqtt"
	"github.com/andig/evcc/util"
)

func init() {
	registry.Add("warp", NewWarpFromConfig)
}

const (
	warpRootTopic = "warp"
	warpTimeout   = 30 * time.Second
)

// NewWarpFromConfig creates a new configurable charger
func NewWarpFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		mqtt.Config `mapstructure:",squash"`
		Topic       string
		Timeout     time.Duration
		UseMeter    bool
	}{
		Topic:   warpRootTopic,
		Timeout: warpTimeout,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	m, err := NewWarp(cc.Config, cc.Topic, cc.Timeout)
	if err != nil {
		return nil, err
	}

	var currentPower func() (float64, error)
	var totalEnergy func() (float64, error)
	if cc.UseMeter {
		currentPower = m.currentPower
		totalEnergy = m.totalEnergy
	}

	return decorateWarp(m, currentPower, totalEnergy), err
}

// Warp configures generic charger and charge meter for an Warp loadpoint
type Warp struct {
	log         *util.Logger
	root        string
	client      *mqtt.Client
	enabledG    func() (string, error)
	statusG     func() (string, error)
	meterG      func() (string, error)
	maxcurrentS func(int64) error
}

//go:generate go run ../cmd/tools/decorate.go -p charger -f decorateWarp -o warp_decorators -b *Warp -r api.Charger -t "api.Meter,CurrentPower,func() (float64, error)" -t "api.MeterEnergy,TotalEnergy,func() (float64, error)"

// NewWarp creates a new configurable charger
func NewWarp(mqttconf mqtt.Config, topic string, timeout time.Duration) (*Warp, error) {
	log := util.NewLogger("warp")

	client, err := mqtt.RegisteredClientOrDefault(log, mqttconf)
	if err != nil {
		return nil, err
	}

	m := &Warp{
		log:    log,
		root:   topic,
		client: client,
	}

	// timeout handler
	timer := provider.NewMqtt(log, client,
		fmt.Sprintf("%s/evse/state", topic), "", 1, timeout,
	).StringGetter()

	stringG := func(topic string) func() (string, error) {
		g := provider.NewMqtt(log, client, topic, "", 1, 0).StringGetter()
		return func() (val string, err error) {
			if val, err = g(); err == nil {
				_, err = timer()
			}
			return val, err
		}
	}

	m.enabledG = stringG(fmt.Sprintf("%s/evse/auto_start_charging", topic))
	m.statusG = stringG(fmt.Sprintf("%s/evse/state", topic))
	m.meterG = stringG(fmt.Sprintf("%s/meter/state", topic))

	m.maxcurrentS = provider.NewMqtt(log, client,
		fmt.Sprintf("%s/evse/current_limit", topic),
		`{ "current": ${maxcurrent} }`, 1, 0,
	).IntSetter("maxcurrent")

	return m, nil
}

type warpStatus struct {
	Iec61851State          int64 `json:"iec61851_state"`
	VehicleState           int64 `json:"vehicle_state"`
	ChargeRelease          int64 `json:"charge_release"`
	ContactorState         int64 `json:"contactor_state"`
	ContactorError         int64 `json:"contactor_error"`
	AllowedChargingCurrent int64 `json:"allowed_charging_current"`
	ErrorState             int64 `json:"error_state"`
	LockState              int64 `json:"lock_state"`
	TimeSinceStateChange   int64 `json:"time_since_state_change"`
	Uptime                 int64 `json:"uptime"`
}

func (m *Warp) status() (warpStatus, error) {
	var res warpStatus

	s, err := m.statusG()
	if err == nil {
		err = json.Unmarshal([]byte(s), &res)
	}

	return res, err
}

// Enable implements the api.Charger interface
func (m *Warp) Enable(enable bool) error {
	action := "stop_charging"
	if enable {
		action = "start_charging"

		// ensure that charger can be enabled
		res, err := m.status()
		if err != nil {
			return err
		}

		if res.ChargeRelease == 2 {
			return errors.New("charger disabled by button or key")
		}
	} else {
		var autostart struct {
			AutoStartCharging bool `json:"auto_start_charging"`
		}

		s, err := m.enabledG()
		if err == nil {
			err = json.Unmarshal([]byte(s), &autostart)
		}

		if err == nil && autostart.AutoStartCharging {
			m.log.WARN.Println("auto_start_charging must be disabled")

			topic := fmt.Sprintf("%s/evse/auto_start_charging_update", m.root)
			if err := m.client.Publish(topic, true, `{ "auto_start_charging": false }`); err != nil {
				return err
			}
		}
	}

	topic := fmt.Sprintf("%s/%s/%s", m.root, "evse", action)

	return m.client.Publish(topic, true, "null")
}

// Enabled implements the api.Charger interface
func (m *Warp) Enabled() (bool, error) {
	res, err := m.status()
	return res.ChargeRelease == 0, err
}

// Status implements the api.Charger interface
func (m *Warp) Status() (api.ChargeStatus, error) {
	var status warpStatus

	s, err := m.statusG()
	if err == nil {
		err = json.Unmarshal([]byte(s), &status)
	}

	res := api.StatusNone
	switch status.VehicleState {
	case 0:
		res = api.StatusA
	case 1:
		res = api.StatusB
	case 2:
		res = api.StatusC
	default:
		if err == nil {
			err = fmt.Errorf("invalid status: %d", status.VehicleState)
		}
	}

	return res, err
}

// MaxCurrent implements the api.Charger interface
func (m *Warp) MaxCurrent(current int64) error {
	return m.maxcurrentS(1000 * current)
}

// MaxCurrentMillis implements the api.ChargerEx interface
func (m *Warp) MaxCurrentMillis(current float64) error {
	return m.maxcurrentS(int64(1000 * current))
}

type powerStatus struct {
	Power     float64 `json:"power"`
	EnergyRel float64 `json:"energy_rel"`
	EnergyAbs float64 `json:"energy_abs"`
}

// currentPower implements the Meter.CurrentPower interface
func (m *Warp) currentPower() (float64, error) {
	var res powerStatus

	s, err := m.meterG()
	if err == nil {
		err = json.Unmarshal([]byte(s), &res)
	}

	return res.Power, err
}

// totalEnergy implements the Meter.TotalEnergy interface
func (m *Warp) totalEnergy() (float64, error) {
	var res powerStatus

	s, err := m.meterG()
	if err == nil {
		err = json.Unmarshal([]byte(s), &res)
	}

	return res.EnergyAbs, err
}
