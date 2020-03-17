package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/charger"
	"github.com/andig/evcc/core"
	"github.com/andig/evcc/core/wrapper"
	"github.com/andig/evcc/provider"
	"github.com/andig/evcc/push"
	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// MQTT singleton
var mq *provider.MqttClient

func clientID() string {
	pid := rand.Int31()
	return fmt.Sprintf("evcc-%d", pid)
}

func configureMeters(conf config) (meters map[string]api.Meter) {
	meters = make(map[string]api.Meter)
	for _, mc := range conf.Meters {
		m := core.NewMeter(floatGetter(mc.Power))

		// decorate Meter with MeterEnergy
		if mc.Energy != nil {
			m = &wrapper.CompositeMeter{
				Meter:       m,
				MeterEnergy: core.NewMeterEnergy(floatGetter(mc.Energy)),
			}
		}
		meters[mc.Name] = m
	}
	return
}

func configureChargers(conf config) (chargers map[string]api.Charger) {
	chargers = make(map[string]api.Charger)
	for _, cc := range conf.Chargers {
		var c api.Charger

		switch cc.Type {
		case "wallbe":
			c = charger.NewWallbe(cc.URI)

		case "default", "configurable":
			c = charger.NewCharger(
				stringGetter(cc.Status),
				boolGetter(cc.Enabled),
				boolSetter("enable", cc.Enable),
				intSetter("current", cc.MaxCurrent),
			)

		default:
			log.FATAL.Fatalf("invalid charger type '%s'", cc.Type)
		}

		chargers[cc.Name] = c
	}
	return
}

func configureSoCs(conf config) (socs map[string]api.SoC) {
	socs = make(map[string]api.SoC)
	for _, sc := range conf.SoCs {
		soc := core.NewSoC(sc.Capacity, sc.Title, floatGetter(sc.Charge))
		socs[sc.Name] = soc
	}
	return
}

// TODO rewrite based on https://github.com/spf13/viper/pull/863
func configureLoadPoint(lp *core.LoadPoint, lpc loadPointConfig, subv *viper.Viper) {
	// for _, key := range []string{"charger", "gridmeter", "pvmeter", "chargemeter", "soc"} {
	// 	delete(kv, key)
	// }

	// config := &mapstructure.DecoderConfig{
	// 	WeaklyTypedInput: true,
	// 	DecodeHook:       mapstructure.StringToTimeDurationHookFunc(),
	// 	Result:           lp,
	// }

	// decoder, err := mapstructure.NewDecoder(config)
	// if err != nil {
	// 	log.FATAL.Fatalf("configuring loadpoints failed: %v", err)
	// }

	// if err := decoder.Decode(kv); err != nil {
	// 	log.FATAL.Fatalf("configuring loadpoints failed: %v", err)
	// }

	// we can ignore the error here as UnmarshalExact has been called before
	_ = subv.UnmarshalExact(lp)

	if lpc.Mode != "" {
		// workaround for golangs yaml off=0 conversion
		if lpc.Mode == "0" {
			lpc.Mode = api.ModeOff
		}
		lp.Mode = lpc.Mode // don't use SetMode here as that will block on channel send
	}
}

func loadConfig(conf config, eventsChan chan push.Event) (loadPoints []*core.LoadPoint) {
	if viper.Get("mqtt") != nil {
		mq = provider.NewMqttClient(conf.Mqtt.Broker, conf.Mqtt.User, conf.Mqtt.Password, clientID(), 1)
	}

	meters := configureMeters(conf)
	chargers := configureChargers(conf)
	socs := configureSoCs(conf)

	for idx, lpc := range conf.LoadPoints {
		// configure loadpoint
		lp := core.NewLoadPoint()
		subv := viper.SubSlice("loadpoints")[idx]
		configureLoadPoint(lp, lpc, subv)

		// assign charger
		if charger, ok := chargers[lpc.Charger]; ok {
			lp.Charger = charger
		} else {
			log.FATAL.Fatalf("invalid charger '%s'", lpc.Charger)
		}

		// assign meters
		for _, m := range []struct {
			key   string
			meter *api.Meter
		}{
			{lpc.GridMeter, &lp.GridMeter},
			{lpc.ChargeMeter, &lp.ChargeMeter},
			{lpc.PVMeter, &lp.PVMeter},
		} {
			if m.key != "" {
				if impl, ok := meters[m.key]; ok {
					*m.meter = impl
				} else {
					log.FATAL.Fatalf("invalid meter '%s'", m.key)
				}
			}
		}

		// assign socs
		if lpc.SoC != "" {
			if impl, ok := socs[lpc.SoC]; ok {
				lp.SoC = impl
			} else {
				log.FATAL.Fatalf("invalid soc '%s'", lpc.SoC)
			}
		}

		loadPoints = append(loadPoints, lp)
	}

	return
}
