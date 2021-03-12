package charger

// Code generated by github.com/andig/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/andig/evcc/api"
)

func decoratecFosPowerBrain(base *cFosPowerBrain, meter func() (float64, error), meterEnergy func() (float64, error), meterCurrent func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterCurrent == nil && meterEnergy == nil:
		return base

	case meter != nil && meterCurrent == nil && meterEnergy == nil:
		return &struct {
			*cFosPowerBrain
			api.Meter
		}{
			cFosPowerBrain: base,
			Meter: &decoratecFosPowerBrainMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			*cFosPowerBrain
			api.MeterEnergy
		}{
			cFosPowerBrain: base,
			MeterEnergy: &decoratecFosPowerBrainMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy != nil:
		return &struct {
			*cFosPowerBrain
			api.Meter
			api.MeterEnergy
		}{
			cFosPowerBrain: base,
			Meter: &decoratecFosPowerBrainMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratecFosPowerBrainMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			*cFosPowerBrain
			api.MeterCurrent
		}{
			cFosPowerBrain: base,
			MeterCurrent: &decoratecFosPowerBrainMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy == nil:
		return &struct {
			*cFosPowerBrain
			api.Meter
			api.MeterCurrent
		}{
			cFosPowerBrain: base,
			Meter: &decoratecFosPowerBrainMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratecFosPowerBrainMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			*cFosPowerBrain
			api.MeterCurrent
			api.MeterEnergy
		}{
			cFosPowerBrain: base,
			MeterCurrent: &decoratecFosPowerBrainMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratecFosPowerBrainMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy != nil:
		return &struct {
			*cFosPowerBrain
			api.Meter
			api.MeterCurrent
			api.MeterEnergy
		}{
			cFosPowerBrain: base,
			Meter: &decoratecFosPowerBrainMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decoratecFosPowerBrainMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decoratecFosPowerBrainMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decoratecFosPowerBrainMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratecFosPowerBrainMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratecFosPowerBrainMeterCurrentImpl struct {
	meterCurrent func() (float64, float64, float64, error)
}

func (impl *decoratecFosPowerBrainMeterCurrentImpl) Currents() (float64, float64, float64, error) {
	return impl.meterCurrent()
}

type decoratecFosPowerBrainMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratecFosPowerBrainMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}
