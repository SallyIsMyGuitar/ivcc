package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateEVSE(base *EVSEWifi, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseVoltages func() (float64, float64, float64, error), chargerEx func(float64) error, identifier func() (string, error)) api.Charger {
	switch {
	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return base

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Meter
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.MeterEnergy
		}{
			EVSEWifi: base,
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy == nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && meterEnergy != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decorateEVSEPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	return nil
}

type decorateEVSEChargerExImpl struct {
	chargerEx func(float64) error
}

func (impl *decorateEVSEChargerExImpl) MaxCurrentMillis(p0 float64) error {
	return impl.chargerEx(p0)
}

type decorateEVSEIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decorateEVSEIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decorateEVSEMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateEVSEMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateEVSEMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateEVSEMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateEVSEPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateEVSEPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decorateEVSEPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decorateEVSEPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}
