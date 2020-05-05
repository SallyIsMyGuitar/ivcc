package core

import (
	"fmt"
	"time"

	"github.com/andig/evcc/api"

	evbus "github.com/asaskevich/EventBus"
	"github.com/benbjohnson/clock"
)

type Ramp struct {
	clock clock.Clock // mockable time
	bus   evbus.Bus   // event bus

	charger api.Charger // Charger

	Name          string
	Sensitivity   int64         // Step size of current change
	MinCurrent    int64         // PV mode: start current	Min+PV mode: min current
	MaxCurrent    int64         // Max allowed current. Physically ensured by the charge controller
	GuardDuration time.Duration // charger enable/disable minimum holding time

	enabled       bool // Charger enabled state
	targetCurrent int64

	// contactor switch guard
	guardUpdated time.Time // charger enabled/disabled timestamp
}

// chargerEnable switches charging on or off. Minimum cycle duration is guaranteed.
func (lp *Ramp) chargerEnable(enable bool) error {
	if lp.targetCurrent != 0 && lp.targetCurrent != lp.MinCurrent {
		log.FATAL.Fatal("charger enable/disable called without setting min current first")
	}

	if remaining := (lp.GuardDuration - time.Since(lp.guardUpdated)).Truncate(time.Second); remaining > 0 {
		log.DEBUG.Printf("%s charger %s - contactor delay %v", lp.Name, status[enable], remaining)
		return nil
	}

	if lp.enabled != enable {
		if err := lp.charger.Enable(enable); err != nil {
			return fmt.Errorf("%s charge controller error: %v", lp.Name, err)
		}

		lp.enabled = enable // cache
		log.INFO.Printf("%s charger %s", lp.Name, status[enable])
		lp.guardUpdated = lp.clock.Now()
	} else {
		log.DEBUG.Printf("%s charger %s", lp.Name, status[enable])
	}

	// if not enabled, current will be reduced to 0 in handler
	lp.bus.Publish(evChargeCurrent, lp.MinCurrent)

	return nil
}

// setTargetCurrent guards setting current against changing to identical value
// and violating MaxCurrent
func (lp *Ramp) setTargetCurrent(targetCurrentIn int64) error {
	targetCurrent := clamp(targetCurrentIn, lp.MinCurrent, lp.MaxCurrent)
	if targetCurrent != targetCurrentIn {
		log.WARN.Printf("%s hard limit charge current: %dA", lp.Name, targetCurrent)
	}

	if lp.targetCurrent != targetCurrent {
		log.DEBUG.Printf("%s set charge current: %dA", lp.Name, targetCurrent)
		if err := lp.charger.MaxCurrent(targetCurrent); err != nil {
			return fmt.Errorf("%s charge controller error: %v", lp.Name, err)
		}

		lp.targetCurrent = targetCurrent // cache
	}

	lp.bus.Publish(evChargeCurrent, targetCurrent)

	return nil
}

// rampUpDown moves stepwise towards target current
func (lp *Ramp) rampUpDown(target int64) error {
	current := lp.targetCurrent
	if current == target {
		return nil
	}

	var step int64
	if current < target {
		step = min(current+lp.Sensitivity, target)
	} else if current > target {
		step = max(current-lp.Sensitivity, target)
	}

	step = clamp(step, lp.MinCurrent, lp.MaxCurrent)

	return lp.setTargetCurrent(step)
}

// rampOff disables charger after setting minCurrent. If already disables, this is a nop.
func (lp *Ramp) rampOff() error {
	if lp.enabled {
		if lp.targetCurrent == lp.MinCurrent {
			return lp.chargerEnable(false)
		}

		return lp.setTargetCurrent(lp.MinCurrent)
	}

	return nil
}

// rampOn enables charger after setting minCurrent. If already enabled, target will be set.
func (lp *Ramp) rampOn(target int64) error {
	if !lp.enabled {
		if err := lp.setTargetCurrent(lp.MinCurrent); err != nil {
			return err
		}

		return lp.chargerEnable(true)
	}

	return lp.setTargetCurrent(target)
}
