package charger

// This file has been generated - do not modify

import (
	"github.com/andig/evcc/api"
)

func decorateEVSE(base api.Charger, meter func() (float64, error), meterEnergy func() (float64, error), meterCurrent func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterCurrent == nil && meterEnergy == nil:
		return base

	case meter != nil && meterCurrent == nil && meterEnergy == nil:
		return &struct{
			api.Charger
			api.Meter
		}{
			Charger: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && meterCurrent == nil && meterEnergy != nil:
		return &struct{
			api.Charger
			api.MeterEnergy
		}{
			Charger: base,
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent == nil && meterEnergy != nil:
		return &struct{
			api.Charger
			api.Meter
			api.MeterEnergy
		}{
			Charger: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy == nil:
		return &struct{
			api.Charger
			api.MeterCurrent
		}{
			Charger: base,
			MeterCurrent: &decorateEVSEMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy == nil:
		return &struct{
			api.Charger
			api.Meter
			api.MeterCurrent
		}{
			Charger: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decorateEVSEMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
		}

	case meter == nil && meterCurrent != nil && meterEnergy != nil:
		return &struct{
			api.Charger
			api.MeterCurrent
			api.MeterEnergy
		}{
			Charger: base,
			MeterCurrent: &decorateEVSEMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterCurrent != nil && meterEnergy != nil:
		return &struct{
			api.Charger
			api.Meter
			api.MeterCurrent
			api.MeterEnergy
		}{
			Charger: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterCurrent: &decorateEVSEMeterCurrentImpl{
				meterCurrent: meterCurrent,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decorateEVSEMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateEVSEMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateEVSEMeterCurrentImpl struct {
	meterCurrent func() (float64, float64, float64, error)
}

func (impl *decorateEVSEMeterCurrentImpl) Currents() (float64, float64, float64, error) {
	return impl.meterCurrent()
}

type decorateEVSEMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateEVSEMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}
