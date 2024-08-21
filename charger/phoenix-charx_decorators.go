package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePhoenixCharx(base *PhoenixCharx, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseVoltages func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return base

	case meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixCharx
			api.Meter
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixCharx
			api.Meter
			api.MeterEnergy
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixCharxMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixCharx
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixCharxPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixCharx
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixCharxMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixCharxPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixCharx
			api.Meter
			api.PhaseVoltages
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decoratePhoenixCharxPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixCharx
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixCharxMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decoratePhoenixCharxPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixCharx
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixCharxPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixCharxPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixCharx
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixCharx: base,
			Meter: &decoratePhoenixCharxMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePhoenixCharxMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decoratePhoenixCharxPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixCharxPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	panic("invalid combination of decorators")
}

type decoratePhoenixCharxMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePhoenixCharxMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePhoenixCharxMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePhoenixCharxMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decoratePhoenixCharxPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixCharxPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decoratePhoenixCharxPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixCharxPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}
