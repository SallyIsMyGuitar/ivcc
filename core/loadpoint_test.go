package core

import (
	"reflect"
	"testing"
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/meter"
	"github.com/andig/evcc/mock"
	"github.com/andig/evcc/provider"
	"github.com/andig/evcc/push"
	"github.com/benbjohnson/clock"
	"github.com/golang/mock/gomock"
)

const (
	lpMinCurrent int64 = 6
	lpMaxCurrent int64 = 16
)

func TestNew(t *testing.T) {
	lp := NewLoadPoint()

	if lp.Mode != api.ModeOff {
		t.Errorf("Mode %v", lp.Mode)
	}
	if lp.Phases != 1 {
		t.Errorf("Phases %v", lp.Phases)
	}
	if lp.MinCurrent != lpMinCurrent {
		t.Errorf("MinCurrent %v", lp.MinCurrent)
	}
	if lp.MaxCurrent != lpMaxCurrent {
		t.Errorf("MaxCurrent %v", lp.MaxCurrent)
	}
	if lp.Sensitivity != 10 {
		t.Errorf("Sensitivity %v", lp.Sensitivity)
	}
	if lp.status != api.StatusNone {
		t.Errorf("status %v", lp.status)
	}
	if lp.enabled {
		t.Errorf("enabled %v", lp.enabled)
	}
	if lp.charging {
		t.Errorf("charging %v", lp.charging)
	}
	if lp.targetCurrent != 0 {
		t.Errorf("targetCurrent %v", lp.targetCurrent)
	}
}

func newLoadPoint(charger api.Charger, pv, gm, cm api.Meter) *LoadPoint {
	lp := NewLoadPoint()
	lp.clock = clock.NewMock()
	lp.clock.(*clock.Mock).Add(time.Hour)

	lp.charger = charger
	lp.pvMeter = pv
	lp.gridMeter = gm

	// prevent assigning a nil pointer sake of
	// https://groups.google.com/forum/#!topic/golang-nuts/wnH302gBa4I/discussion
	if !(cm == nil || reflect.ValueOf(cm).IsNil()) {
		lp.chargeMeter = cm
	}

	uiChan := make(chan Param)
	notificationChan := make(chan push.Event)

	lp.Prepare(uiChan, notificationChan)

	go func() {
		for {
			select {
			case <-uiChan:
			case <-notificationChan:
			}
		}
	}()

	return lp
}

func newEnvironment(t *testing.T, ctrl *gomock.Controller, pm, gm, cm api.Meter) (*LoadPoint, *mock.MockCharger) {
	wb := mock.NewMockCharger(ctrl)

	wb.EXPECT().Enabled().Return(true, nil) // initial alignment with wb
	wb.EXPECT().MaxCurrent(lpMinCurrent)    // initial alignment with wb

	lp := newLoadPoint(wb, pm, gm, cm)
	if !lp.enabled {
		t.Errorf("enabled %v", lp.enabled)
	}
	if lp.guardUpdated != lp.clock.Now() {
		t.Errorf("guardUpdated %v", lp.guardUpdated)
	}

	return lp, wb
}

func TestMeterConfigurations(t *testing.T) {
	tc := []struct {
		gm, cm, pm bool
	}{
		// {false, false, false}, // no meter
		// {false, true, false}, // cm only
		{false, false, true}, // pm only
		{true, false, false}, // gm only
		{true, true, false},  // gm + cm
		{true, false, true},  // gm + pm
		{true, true, true},   // gm + cm + pm
	}

	fg := provider.FloatGetter(func() (float64, error) {
		return 1, nil
	})

	for _, tc := range tc {
		t.Logf("gm: %+v  cm: %v  pm: %v", tc.gm, tc.cm, tc.pm)

		var gm, pm, cm api.Meter
		if tc.gm {
			gm = meter.NewConfigurable(fg)
		}
		if tc.cm {
			cm = meter.NewConfigurable(fg)
		}
		if tc.pm {
			pm = meter.NewConfigurable(fg)
		}

		ctrl := gomock.NewController(t)
		lp, wb := newEnvironment(t, ctrl, pm, gm, cm)
		wb.EXPECT().Status().Return(api.StatusA, nil) // disconnected
		wb.EXPECT().Enable(false)                     // "off" mode

		lp.update()
	}
}

func TestInitialUpdate(t *testing.T) {
	tc := []struct {
		status api.ChargeStatus
		mode   api.ChargeMode
	}{
		{status: api.StatusA, mode: api.ModeOff},
		{status: api.StatusA, mode: api.ModeNow},
		{status: api.StatusA, mode: api.ModeMinPV},
		{status: api.StatusA, mode: api.ModePV},

		{status: api.StatusB, mode: api.ModeOff},
		{status: api.StatusB, mode: api.ModeNow},
		{status: api.StatusB, mode: api.ModeMinPV},
		{status: api.StatusB, mode: api.ModePV},

		{status: api.StatusC, mode: api.ModeOff},
		{status: api.StatusC, mode: api.ModeNow},
		{status: api.StatusC, mode: api.ModeMinPV},
		{status: api.StatusC, mode: api.ModePV},
	}

	for _, tc := range tc {
		t.Logf("%+v\n", tc)

		ctrl := gomock.NewController(t)

		pm := mock.NewMockMeter(ctrl)
		gm := mock.NewMockMeter(ctrl)
		cm := mock.NewMockMeter(ctrl)
		// cm = nil

		lp, wb := newEnvironment(t, ctrl, pm, gm, cm)
		lp.Mode = tc.mode

		wb.EXPECT().Status().Return(tc.status, nil)

		// values are relevant for PV case
		minPower := float64(lpMinCurrent) * lp.Voltage
		pm.EXPECT().CurrentPower().Return(minPower, nil)
		gm.EXPECT().CurrentPower().Return(float64(0), nil)
		if cm != nil {
			cm.EXPECT().CurrentPower().Return(minPower, nil)
		}

		// disable if not connected
		if tc.mode == api.ModeOff {
			wb.EXPECT().Enable(false)
		}

		// power up if now
		if tc.status != api.StatusA && tc.mode == api.ModeNow {
			wb.EXPECT().MaxCurrent(lpMaxCurrent)
		}

		lp.update()

		// max current if connected & mode now
		if tc.status != api.StatusA && tc.mode == api.ModeNow {
			if lp.targetCurrent != lpMaxCurrent {
				t.Errorf("targetCurrent %v", lp.targetCurrent)
			}
		}

		// min current in first cycle
		if tc.mode != api.ModeNow {
			if lp.targetCurrent != lpMinCurrent {
				t.Errorf("targetCurrent %v", lp.targetCurrent)
			}
		}

		// status c means charging
		if lp.charging != (tc.status == api.StatusC) {
			t.Errorf("charging %v", lp.charging)
		}

		ctrl.Finish()
	}
}

func TestImmediateOnOff(t *testing.T) {
	tc := []struct {
		status api.ChargeStatus
		mode   api.ChargeMode
	}{
		{status: api.StatusC, mode: api.ModePV},
	}

	for _, tc := range tc {
		t.Logf("%+v\n", tc)

		ctrl := gomock.NewController(t)

		pm := mock.NewMockMeter(ctrl)
		gm := mock.NewMockMeter(ctrl)
		cm := mock.NewMockMeter(ctrl)
		// cm = nil

		lp, wb := newEnvironment(t, ctrl, pm, gm, cm)
		lp.Mode = tc.mode

		// -- round 1
		wb.EXPECT().Status().Return(tc.status, nil)

		// values are relevant for PV case
		minPower := float64(lpMinCurrent) * lp.Voltage * float64(lp.Phases)
		pm.EXPECT().CurrentPower().Return(minPower, nil)
		gm.EXPECT().CurrentPower().Return(0.0, nil)
		if cm != nil {
			cm.EXPECT().CurrentPower().Return(minPower, nil)
		}

		// disable if not connected
		if tc.status != api.StatusA && tc.mode == api.ModeOff {
			wb.EXPECT().Enable(false)
		}

		// power up if now
		if tc.status != api.StatusA && tc.mode == api.ModeNow {
			wb.EXPECT().MaxCurrent(lpMaxCurrent)
		}

		lp.update()

		// max current if connected & mode now
		if tc.status != api.StatusA && tc.mode == api.ModeNow {
			if lp.targetCurrent != lpMaxCurrent {
				t.Errorf("targetCurrent %v", lp.targetCurrent)
			}
		}

		// min current in first cycle
		if tc.mode != api.ModeNow {
			if lp.targetCurrent != lpMinCurrent {
				t.Errorf("targetCurrent %v", lp.targetCurrent)
			}
		}

		// status c means charging
		if lp.charging != (tc.status == api.StatusC) {
			t.Errorf("charging %v", lp.charging)
		}

		// -- round 2
		wb.EXPECT().Status().Return(tc.status, nil)

		pm.EXPECT().CurrentPower().Return(minPower, nil)
		gm.EXPECT().CurrentPower().Return(-2*minPower, nil)
		if cm != nil {
			cm.EXPECT().CurrentPower().Return(1.0, nil)
		}

		wb.EXPECT().MaxCurrent(2 * lpMinCurrent)

		lp.update()

		// -- round 3
		t.Logf("%+v - 3 (status: %v, enabled: %v, current %d)\n", tc, lp.status, lp.enabled, lp.targetCurrent)

		wb.EXPECT().Status().Return(tc.status, nil)

		pm.EXPECT().CurrentPower().Return(minPower, nil)
		gm.EXPECT().CurrentPower().Return(-2*minPower, nil)
		if cm != nil {
			cm.EXPECT().CurrentPower().Return(1.0, nil)
		}

		wb.EXPECT().MaxCurrent(lpMinCurrent)

		lp.SetMode(api.ModeOff)
		lp.update()

		ctrl.Finish()
	}
}

func TestConsumedPower(t *testing.T) {
	tc := []struct {
		grid, pv, battery, consumed float64
	}{
		{0, 0, 0, 0},    // silent night
		{1, 0, 0, 1},    // grid import
		{0, 1, 0, 1},    // pv sign ignored
		{0, -1, 0, 1},   // pv sign ignored
		{1, 1, 0, 2},    // grid import + pv, pv sign ignored
		{1, -1, 0, 2},   // grid import + pv, pv sign ignored
		{0, 0, 1, 1},    // battery discharging
		{0, 0, -1, -1},  // battery charging -> negative result cannot occur in reality
		{1, -3, 1, 5},   // grid import + pv + battery discharging
		{1, -3, -1, 3},  // grid import + pv + battery charging -> should not happen in reality
		{0, -3, -1, 2},  // pv + battery charging
		{-1, -4, -1, 2}, // grid export + pv + battery charging
		{-1, -4, 0, 3},  // grid export + pv
	}

	for _, tc := range tc {
		res := consumedPower(tc.pv, tc.battery, tc.grid)
		if res != tc.consumed {
			t.Errorf("consumedPower wanted %.f, got %.f", tc.consumed, res)
		}
	}
}
