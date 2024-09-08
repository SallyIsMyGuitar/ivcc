package main

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateTest(base api.Charger, meterEnergy func() (float64, error), phaseSwitcher func(int) error, phaseGetter func() (int, error)) api.Charger {
	switch {
	case meterEnergy == nil && phaseSwitcher == nil:
		return base

	case meterEnergy != nil && phaseSwitcher == nil:
		return &struct {
			api.Charger
			api.MeterEnergy
		}{
			Charger: base,
			MeterEnergy: &decorateTestMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meterEnergy == nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			api.Charger
			api.PhaseSwitcher
		}{
			Charger: base,
			PhaseSwitcher: &decorateTestPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meterEnergy != nil && phaseGetter == nil && phaseSwitcher != nil:
		return &struct {
			api.Charger
			api.MeterEnergy
			api.PhaseSwitcher
		}{
			Charger: base,
			MeterEnergy: &decorateTestMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseSwitcher: &decorateTestPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meterEnergy == nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			api.Charger
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			Charger: base,
			PhaseGetter: &decorateTestPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateTestPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}

	case meterEnergy != nil && phaseGetter != nil && phaseSwitcher != nil:
		return &struct {
			api.Charger
			api.MeterEnergy
			api.PhaseGetter
			api.PhaseSwitcher
		}{
			Charger: base,
			MeterEnergy: &decorateTestMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseGetter: &decorateTestPhaseGetterImpl{
				phaseGetter: phaseGetter,
			},
			PhaseSwitcher: &decorateTestPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	panic("invalid combination of decorators")
}

type decorateTestMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateTestMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateTestPhaseGetterImpl struct {
	phaseGetter func() (int, error)
}

func (impl *decorateTestPhaseGetterImpl) GetPhases() (int, error) {
	return impl.phaseGetter()
}

type decorateTestPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateTestPhaseSwitcherImpl) Phases1p3p(p0 int) error {
	return impl.phaseSwitcher(p0)
}
