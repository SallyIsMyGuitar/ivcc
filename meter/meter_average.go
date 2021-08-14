package meter

import (
	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
)

func init() {
	registry.Add("movingaverage", NewMovingAverageFromConfig)
}

// NewMovingAverageFromConfig creates api.Meter from config
func NewMovingAverageFromConfig(other map[string]interface{}) (api.Meter, error) {
	cc := struct {
		Decay float64
		Meter struct {
			Type  string
			Other map[string]interface{} `mapstructure:",remain"`
		}
	}{
		Decay: 0.1,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	m, err := NewFromConfig(cc.Meter.Type, cc.Meter.Other)
	if err != nil {
		return nil, err
	}

	mav := &MovingAverage{
		decay:         cc.Decay,
		currentPowerG: m.CurrentPower,
	}

	meter, _ := NewConfigurable(mav.CurrentPower)

	// decorate energy reading
	var totalEnergy func() (float64, error)
	if m, ok := m.(api.MeterEnergy); ok {
		totalEnergy = m.TotalEnergy
	}

	// decorate battery reading
	var batterySoC func() (float64, error)
	if m, ok := m.(api.Battery); ok {
		batterySoC = m.SoC
	}

	// decorate currents reading
	var currents func() (float64, float64, float64, error)
	if m, ok := m.(api.MeterCurrent); ok {
		currents = m.Currents
	}

	res := meter.Decorate(totalEnergy, currents, batterySoC)

	return res, nil
}

type MovingAverage struct {
	decay         float64
	value         *float64
	currentPowerG func() (float64, error)
}

func (m *MovingAverage) CurrentPower() (float64, error) {
	power, err := m.currentPowerG()
	if err != nil {
		return power, err
	}

	m.add(power)

	return m.get(), nil
}

// modeled after https://github.com/VividCortex/ewma

// Add adds a value to the series and updates the moving average.
func (m *MovingAverage) add(value float64) {
	if m.value == nil { // this is a proxy for "uninitialized"
		m.value = &value
	} else {
		*m.value = (value * m.decay) + (m.get() * (1 - m.decay))
	}
}

// Value returns the current value of the moving averagm.
func (m *MovingAverage) get() float64 {
	if m.value == nil { // this is a proxy for "uninitialized"
		return 0
	} else {
		return *m.value
	}
}
