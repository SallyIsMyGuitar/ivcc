package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateMeter(base api.Meter, meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseVoltages func() (float64, float64, float64, error), phasePowers func() (float64, float64, float64, error), battery func() (float64, error), batteryCapacity func() float64, batteryController func(api.BatteryMode) error) api.Meter {
	switch {
	case battery == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return base

	case battery == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
		}{
			Meter: base,
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
		}{
			Meter: base,
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController == nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers == nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy == nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && batteryController != nil && meterEnergy != nil && phaseCurrents != nil && phasePowers != nil && phaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.BatteryController
			api.MeterEnergy
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateMeterBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateMeterBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			BatteryController: &decorateMeterBatteryControllerImpl{
				batteryController: batteryController,
			},
			MeterEnergy: &decorateMeterMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateMeterPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhasePowers: &decorateMeterPhasePowersImpl{
				phasePowers: phasePowers,
			},
			PhaseVoltages: &decorateMeterPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	panic("invalid combination of decorators")
}

type decorateMeterBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateMeterBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateMeterBatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decorateMeterBatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}

type decorateMeterBatteryControllerImpl struct {
	batteryController func(api.BatteryMode) error
}

func (impl *decorateMeterBatteryControllerImpl) SetBatteryMode(p0 api.BatteryMode) error {
	return impl.batteryController(p0)
}

type decorateMeterMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateMeterMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateMeterPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateMeterPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decorateMeterPhasePowersImpl struct {
	phasePowers func() (float64, float64, float64, error)
}

func (impl *decorateMeterPhasePowersImpl) Powers() (float64, float64, float64, error) {
	return impl.phasePowers()
}

type decorateMeterPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decorateMeterPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}
