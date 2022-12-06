package id

import (
	"fmt"
	"strings"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/provider"
)

// Provider is an api.Vehicle implementation for VW ID cars
type Provider struct {
	statusG func() (SelectiveSatus, error)
	action  func(action, value string) error
}

// NewProvider creates a new vehicle
func NewProvider(api *API, vin string, cache time.Duration) *Provider {
	impl := &Provider{
		statusG: provider.Cached(func() (SelectiveSatus, error) {
			return api.Status(vin)
		}, cache),
		action: func(action, value string) error {
			return api.Action(vin, action, value)
		},
	}
	return impl
}

var _ api.Battery = (*Provider)(nil)

// SoC implements the api.Vehicle interface
func (v *Provider) SoC() (float64, error) {
	res, err := v.statusG()
	if err != nil {
		return 0, err
	}

	if res.Charging == nil {
		return 0, fmt.Errorf("SoC not avaliable")
	}

	return float64(res.Charging.BatteryStatus.Value.CurrentSOCPct), nil
}

var _ api.ChargeState = (*Provider)(nil)

// Status implements the api.ChargeState interface
func (v *Provider) Status() (api.ChargeStatus, error) {
	status := api.StatusA // disconnected

	res, err := v.statusG()
	if err != nil {
		return "", err
	}

	if res.Charging == nil {
		return "", fmt.Errorf("PlugStatus not avaliable")
	}

	if res.Charging.PlugStatus.Value.PlugConnectionState == "connected" {
		status = api.StatusB
	}
	if res.Charging.ChargingStatus.Value.ChargingState == "charging" {
		status = api.StatusC
	}

	return status, err
}

var _ api.VehicleFinishTimer = (*Provider)(nil)

// FinishTime implements the api.VehicleFinishTimer interface
func (v *Provider) FinishTime() (time.Time, error) {
	res, err := v.statusG()
	if err != nil {
		return time.Time{}, err
	}

	if res.Charging == nil {
		return time.Time{}, fmt.Errorf("Finishtime not avaliable")
	}

	return res.Charging.ChargingStatus.Value.CarCapturedTimestamp.Add(time.Duration(res.Charging.ChargingStatus.Value.RemainingChargingTimeToCompleteMin) * time.Minute), err

}

var _ api.VehicleRange = (*Provider)(nil)

// Range implements the api.VehicleRange interface
func (v *Provider) Range() (int64, error) {
	res, err := v.statusG()
	if err != nil {
		return 0, err
	}

	if res.Charging == nil {
		return 0, fmt.Errorf("Range not avaliable")
	}

	return int64(res.Charging.BatteryStatus.Value.CruisingRangeElectricKm), nil

}

var _ api.VehicleOdometer = (*Provider)(nil)

// Odometer implements the api.VehicleOdometer interface
func (v *Provider) Odometer() (float64, error) {
	return 0, fmt.Errorf("Odometer not avaliable")
}

var _ api.VehicleClimater = (*Provider)(nil)

// Climater implements the api.VehicleClimater interface
func (v *Provider) Climater() (active bool, outsideTemp, targetTemp float64, err error) {
	res, err := v.statusG()
	if err != nil {
		return active, outsideTemp, targetTemp, err
	}

	if res.Climatisation == nil {
		return false, 0, 0, fmt.Errorf("Climater not avaliable")
	}

	state := strings.ToLower(res.Climatisation.ClimatisationStatus.Value.ClimatisationState)

	if state == "" {
		return false, 0, 0, api.ErrNotAvailable
	}

	active = state != "off" && state != "invalid" && state != "error"
	targetTemp = float64(res.Climatisation.ClimatisationSettings.Value.TargetTemperatureC)
	// TODO not available; use target temp to avoid wrong heating/cooling display
	outsideTemp = targetTemp

	return active, outsideTemp, targetTemp, nil

}

var _ api.SocLimiter = (*Provider)(nil)

// TargetSoC implements the api.SocLimiter interface
func (v *Provider) TargetSoC() (float64, error) {
	res, err := v.statusG()
	if err != nil {
		return 0, err
	}

	if res.Charging == nil {
		return 0, fmt.Errorf("Target SoC not avaliable")
	}

	return float64(res.Charging.ChargingSettings.Value.TargetSOCPct), nil
}

var _ api.VehicleChargeController = (*Provider)(nil)

// StartCharge implements the api.VehicleChargeController interface
func (v *Provider) StartCharge() error {
	return v.action(ActionCharge, ActionChargeStart)
}

// StopCharge implements the api.VehicleChargeController interface
func (v *Provider) StopCharge() error {
	return v.action(ActionCharge, ActionChargeStop)
}
