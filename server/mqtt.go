package server

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/loadpoint"
	"github.com/evcc-io/evcc/core/site"
	"github.com/evcc-io/evcc/provider/mqtt"
	"github.com/evcc-io/evcc/util"
)

var deprecatedTopics = []string{
	"activePhases", "range", "socCharge",
	"vehicleSoC", "batterySoC", "bufferSoC", "minSoC", "prioritySoC", "targetSoC", "vehicleTargetSoC",
	"savingsAmount", "savingsEffectivePrice", "savingsGridCharged", "savingsSelfConsumptionCharged", "savingsSelfConsumptionPercent", "savingsTotalCharged",
	"stats/30d", "stats/365d", "stats/total",
}

// MQTT is the MQTT server. It uses the MQTT client for publishing.
type MQTT struct {
	log     *util.Logger
	Handler *mqtt.Client
	root    string
}

// NewMQTT creates MQTT server
func NewMQTT(root string) *MQTT {
	return &MQTT{
		log:     util.NewLogger("mqtt"),
		Handler: mqtt.Instance,
		root:    root,
	}
}

func (m *MQTT) encode(v interface{}) string {
	// nil should erase the value
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case float64:
		return fmt.Sprintf("%.5g", val)
	case time.Time:
		if val.IsZero() {
			return ""
		}
		return strconv.FormatInt(val.Unix(), 10)
	case time.Duration:
		// must be before stringer to convert to seconds instead of string
		return fmt.Sprintf("%d", int64(val.Seconds()))
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}

func (m *MQTT) publishComplex(topic string, retained bool, payload interface{}) {
	if payload == nil {
		m.publishSingleValue(topic, retained, payload)
		return
	}

	switch typ := reflect.TypeOf(payload); typ.Kind() {
	case reflect.Slice:
		// publish count
		val := reflect.ValueOf(payload)
		m.publishSingleValue(topic, retained, val.Len())

		// loop slice
		for i := 0; i < val.Len(); i++ {
			m.publishComplex(fmt.Sprintf("%s/%d", topic, i+1), retained, val.Index(i).Interface())
		}

	case reflect.Map:
		// loop map
		for iter := reflect.ValueOf(payload).MapRange(); iter.Next(); {
			k := iter.Key().String()
			m.publishComplex(fmt.Sprintf("%s/%s", topic, k), retained, iter.Value().Interface())
		}

	case reflect.Struct:
		val := reflect.ValueOf(payload)
		typ := val.Type()

		// loop struct
		for i := 0; i < typ.NumField(); i++ {
			if f := typ.Field(i); f.IsExported() {
				n := f.Name
				m.publishComplex(fmt.Sprintf("%s/%s", topic, strings.ToLower(n[:1])+n[1:]), retained, val.Field(i).Interface())
			}
		}

	default:
		m.publishSingleValue(topic, retained, payload)
	}
}

func (m *MQTT) publishSingleValue(topic string, retained bool, payload interface{}) {
	token := m.Handler.Client.Publish(topic, m.Handler.Qos, retained, m.encode(payload))
	go m.Handler.WaitForToken("send", topic, token)
}

func (m *MQTT) publish(topic string, retained bool, payload interface{}) {
	// publish phase values
	if slice, ok := payload.([]float64); ok && len(slice) == 3 {
		var total float64
		for i, v := range slice {
			total += v
			m.publishSingleValue(fmt.Sprintf("%s/l%d", topic, i+1), retained, v)
		}

		// publish sum value
		m.publishSingleValue(topic, retained, total)

		return
	}

	m.publishComplex(topic, retained, payload)
}

// TODO add plan api
// func (m *MQTT) listenVehicleSetters(topic string, site site.API, v vehicle.API) {
// 	m.Handler.ListenSetter(topic+"/minSoc", func(payload string) error {
// 		soc, err := strconv.Atoi(payload)
// 		if err == nil {
// 			v.SetMinSoc(soc)
// 		}
// 		return err
// 	})
// 	m.Handler.ListenSetter(topic+"/limitSoc", func(payload string) error {
// 		soc, err := strconv.Atoi(payload)
// 		if err == nil {
// 			v.SetLimitSoc(soc)
// 		}
// 		return err
// 	})
// 	// m.Handler.ListenSetter(topic+"/mode", func(payload string) error {
// 	// 	mode, err := api.ChargeModeString(payload)
// 	// 	if err == nil {
// 	// 		v.SetMode(mode)
// 	// 	}
// 	// 	return err
// 	// })
// 	// m.Handler.ListenSetter(topic+"/phases", func(payload string) error {
// 	// 	phases, err := strconv.Atoi(payload)
// 	// 	if err == nil {
// 	// 		err = v.SetPhases(phases)
// 	// 	}
// 	// 	return err
// 	// })
// 	// m.Handler.ListenSetter(topic+"/minCurrent", func(payload string) error {
// 	// 	current, err := parseFloat(payload)
// 	// 	if err == nil {
// 	// 		v.SetMinCurrent(current)
// 	// 	}
// 	// 	return err
// 	// })
// 	// m.Handler.ListenSetter(topic+"/maxCurrent", func(payload string) error {
// 	// 	current, err := parseFloat(payload)
// 	// 	if err == nil {
// 	// 		v.SetMaxCurrent(current)
// 	// 	}
// 	// 	return err
// 	// })
// }

<<<<<<< HEAD
func (m *MQTT) listenLoadpointSetters(topic string, site site.API, lp loadpoint.API) error {
	var err error

	if err == nil {
		err = m.Handler.ListenSetter(topic+"/mode", func(payload string) error {
			mode, err := api.ChargeModeString(payload)
			if err == nil {
				lp.SetMode(mode)
=======
func (m *MQTT) listenLoadpointSetters(topic string, site site.API, lp loadpoint.API) {
	m.Handler.ListenSetter(topic+"/mode", func(payload string) error {
		mode, err := api.ChargeModeString(payload)
		if err == nil {
			lp.SetMode(mode)
		}
		return err
	})
	m.Handler.ListenSetter(topic+"/limitSoc", func(payload string) error {
		soc, err := strconv.Atoi(payload)
		if err == nil {
			lp.SetLimitSoc(soc)
		}
		return err
	})
	// TODO plan
	// TODO vehicle
	// m.Handler.ListenSetter(topic+"/planEnergy", func(payload string) error {
	// 	val, err := parseFloat(payload)
	// 	if err == nil {
	// 		lp.SetPlanEnergy(val)
	// 	}
	// 	return err
	// })
	// m.Handler.ListenSetter(topic+"/planSoc", func(payload string) error {
	// 	soc, err := strconv.Atoi(payload)
	// 	if err == nil {
	// 		lp.SetPlanSoc(soc)
	// 	}
	// 	return err
	// })
	// m.Handler.ListenSetter(topic+"/planTime", func(payload string) error {
	// 	val, err := time.Parse(time.RFC3339, payload)
	// 	if err == nil {
	// 		err = lp.SetPlanTime(val)
	// 	} else if string(payload) == "null" {
	// 		err = lp.SetPlanTime(time.Time{})
	// 	}
	// 	return err
	// })
	m.Handler.ListenSetter(topic+"/minCurrent", func(payload string) error {
		current, err := parseFloat(payload)
		if err == nil {
			lp.SetMinCurrent(current)
		}
		return err
	})
	m.Handler.ListenSetter(topic+"/maxCurrent", func(payload string) error {
		current, err := parseFloat(payload)
		if err == nil {
			lp.SetMaxCurrent(current)
		}
		return err
	})
	m.Handler.ListenSetter(topic+"/phases", func(payload string) error {
		phases, err := strconv.Atoi(payload)
		if err == nil {
			err = lp.SetPhases(phases)
		}
		return err
	})
	m.Handler.ListenSetter(topic+"/vehicle", func(payload string) error {
		vehicle, err := strconv.Atoi(payload)
		if err == nil {
			if vehicle > 0 {
				if vehicles := site.Vehicles().All(); vehicle <= len(vehicles) {
					lp.SetVehicle(vehicles[vehicle-1])
				} else {
					err = fmt.Errorf("invalid vehicle: %d", vehicle)
				}
			} else {
				lp.SetVehicle(nil)
>>>>>>> 20de10caf (Add vehicle settings api)
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/limitSoc", func(payload string) error {
			soc, err := strconv.Atoi(payload)
			if err == nil {
				lp.SetSessionSocLimit(soc)
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/planEnergy", func(payload string) error {
			val, err := parseFloat(payload)
			if err == nil {
				lp.SetPlanEnergy(val)
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/planSoc", func(payload string) error {
			soc, err := strconv.Atoi(payload)
			if err == nil {
				lp.SetPlanSoc(soc)
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/planTime", func(payload string) error {
			val, err := time.Parse(time.RFC3339, payload)
			if err == nil {
				err = lp.SetPlanTime(val)
			} else if string(payload) == "null" {
				err = lp.SetPlanTime(time.Time{})
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/minCurrent", func(payload string) error {
			current, err := parseFloat(payload)
			if err == nil {
				lp.SetMinCurrent(current)
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/maxCurrent", func(payload string) error {
			current, err := parseFloat(payload)
			if err == nil {
				lp.SetMaxCurrent(current)
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/phases", func(payload string) error {
			phases, err := strconv.Atoi(payload)
			if err == nil {
				err = lp.SetPhases(phases)
			}
			return err
		})
	}
	if err == nil {
		// TODO name
		err = m.Handler.ListenSetter(topic+"/vehicle", func(payload string) error {
			vehicle, err := strconv.Atoi(payload)
			if err == nil {
				if vehicle > 0 {
					if vehicles := site.GetVehicles(); vehicle <= len(vehicles) {
						lp.SetVehicle(vehicles[vehicle-1])
					} else {
						err = fmt.Errorf("invalid vehicle: %d", vehicle)
					}
				} else {
					lp.SetVehicle(nil)
				}
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/enableThreshold", func(payload string) error {
			threshold, err := parseFloat(payload)
			if err == nil {
				lp.SetEnableThreshold(threshold)
			}
			return err
		})
	}
	if err == nil {
		err = m.Handler.ListenSetter(topic+"/disableThreshold", func(payload string) error {
			threshold, err := parseFloat(payload)
			if err == nil {
				lp.SetDisableThreshold(threshold)
			}
			return err
		})
	}

	return err
}

// Run starts the MQTT publisher for the MQTT API
func (m *MQTT) Run(site site.API, in <-chan util.Param) {
	// site setters
	if err := m.Handler.ListenSetter(m.root+"/site/prioritySoc", func(payload string) error {
		val, err := parseFloat(payload)
		if err == nil {
			err = site.SetPrioritySoc(val)
		}
		return err
	}); err != nil {
		m.log.ERROR.Println(err)
	}

	if err := m.Handler.ListenSetter(m.root+"/site/bufferSoc", func(payload string) error {
		val, err := parseFloat(payload)
		if err == nil {
			err = site.SetBufferSoc(val)
		}
		return err
	}); err != nil {
		m.log.ERROR.Println(err)
	}

	if err := m.Handler.ListenSetter(m.root+"/site/bufferStartSoc", func(payload string) error {
		val, err := parseFloat(payload)
		if err == nil {
			err = site.SetBufferStartSoc(val)
		}
		return err
	}); err != nil {
		m.log.ERROR.Println(err)
	}

	if err := m.Handler.ListenSetter(m.root+"/site/residualPower", func(payload string) error {
		val, err := parseFloat(payload)
		if err == nil {
			err = site.SetResidualPower(val)
		}
		return err
	}); err != nil {
		m.log.ERROR.Println(err)
	}

	if err := m.Handler.ListenSetter(m.root+"/site/smartCostLimit", func(payload string) error {
		val, err := parseFloat(payload)
		if err == nil {
			err = site.SetSmartCostLimit(val)
		}
		return err
	}); err != nil {
		m.log.ERROR.Println(err)
	}

	// number of vehicles
	topic := fmt.Sprintf("%s/vehicles", m.root)
	m.publish(topic, true, len(site.Vehicles().All()))

	// TODO decide key
	// vehicle setters
	// for id, vehicle := range site.Vehicles().All() {
	// 	topic := fmt.Sprintf("%s/vehicles/%d", m.root, id+1)
	// 	m.listenVehicleSetters(topic, site, vehicle)
	// }

	// number of loadpoints
	topic = fmt.Sprintf("%s/loadpoints", m.root)
	m.publish(topic, true, len(site.Loadpoints()))

	// loadpoint setters
	for id, lp := range site.Loadpoints() {
		topic := fmt.Sprintf("%s/loadpoints/%d", m.root, id+1)
		if err := m.listenLoadpointSetters(topic, site, lp); err != nil {
			m.log.ERROR.Println(err)
		}
	}

	// TODO remove deprecated topics
	for _, dep := range deprecatedTopics {
		m.publish(fmt.Sprintf("%s/site/%s", m.root, dep), true, nil)
	}

	for id := range site.Loadpoints() {
		topic := fmt.Sprintf("%s/loadpoints/%d", m.root, id+1)
		for _, dep := range deprecatedTopics {
			m.publish(fmt.Sprintf("%s/%s", topic, dep), true, nil)
		}
	}

	for i := 0; i < 10; i++ {
		m.publish(fmt.Sprintf("%s/site/pv/%d", m.root, i), true, nil)
		m.publish(fmt.Sprintf("%s/site/battery/%d", m.root, i), true, nil)
		m.publish(fmt.Sprintf("%s/site/vehicles/%d", m.root, i), true, nil)
	}

	// alive indicator
	var updated time.Time

	// publish
	for p := range in {
		topic := fmt.Sprintf("%s/site", m.root)
		if p.Loadpoint != nil {
			id := *p.Loadpoint + 1
			topic = fmt.Sprintf("%s/loadpoints/%d", m.root, id)
		}

		// alive indicator
		if time.Since(updated) > time.Second {
			updated = time.Now()
			m.publish(fmt.Sprintf("%s/updated", m.root), true, updated.Unix())
		}

		// value
		topic += "/" + p.Key
		m.publish(topic, true, p.Val)
	}
}

// parseFloat rejects NaN and Inf values
func parseFloat(payload string) (float64, error) {
	f, err := strconv.ParseFloat(payload, 64)
	if err == nil && (math.IsNaN(f) || math.IsInf(f, 0)) {
		err = fmt.Errorf("invalid float value: %s", payload)
	}
	return f, err
}
