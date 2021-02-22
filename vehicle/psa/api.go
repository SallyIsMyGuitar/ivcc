package psa

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/andig/evcc/util"
	"github.com/andig/evcc/util/request"
)

// https://github.com/flobz/psa_car_controller
// https://github.com/snaptec/openWB/blob/master/modules/soc_psa/psasoc.py

// API is an api.Vehicle implementation for PSA cars
type API struct {
	*request.Helper
	brand, realm, id, secret string
}

// NewAPI creates a new vehicle
func NewAPI(log *util.Logger, brand, realm, id, secret string) *API {
	v := &API{
		Helper: request.NewHelper(log),
		brand:  brand,
		realm:  realm,
		id:     id,
		secret: secret,
	}
	return v
}

func (v *API) Login(user, password string) error {
	data := url.Values{
		"realm":      []string{v.realm},
		"scope":      []string{"openid profile"},
		"grant_type": []string{"password"},
		"username":   []string{user},
		"password":   []string{password},
	}

	uri := fmt.Sprintf("https://idpcvs.%s/am/oauth2/access_token", v.brand)
	req, err := request.New(http.MethodPost, uri, strings.NewReader(data.Encode()), map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": "Basic %s",
	})

	var resp http.Response
	if err == nil {
		resp, err = v.Do(req)
	}
	_ = resp

	return err
}

// Vehicles implements the /vehicles response
func (v *API) Vehicles() (res []string, err error) {
	uri := "https://mobileapi.apps.emea.vwapps.io/vehicles"

	req, err := request.New(http.MethodGet, uri, nil, map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + "v.identity.Token()",
	})

	var vehicles struct {
		Data []struct {
			VIN      string
			Nickname string
		}
	}

	if err == nil {
		err = v.DoJSON(req, &vehicles)

		for _, v := range vehicles.Data {
			res = append(res, v.VIN)
		}
	}

	return res, err
}

// Status is the /status api
type Status struct {
	Data struct {
		BatteryStatus         BatteryStatus
		ChargingStatus        ChargingStatus
		ChargingSettings      ChargingSettings
		PlugStatus            PlugStatus
		RangeStatus           RangeStatus
		ClimatisationSettings ClimatisationSettings
		ClimatisationStatus   ClimatisationStatus // may be currently not available
	}
}

// BatteryStatus is the /status.batteryStatus api
type BatteryStatus struct {
	CarCapturedTimestamp    string
	CurrentSOCPercent       int `json:"currentSOC_pct"`
	CruisingRangeElectricKm int `json:"cruisingRangeElectric_km"`
}

// ChargingStatus is the /status.chargingStatus api
type ChargingStatus struct {
	CarCapturedTimestamp               string
	ChargingState                      string // readyForCharging
	RemainingChargingTimeToCompleteMin int    `json:"remainingChargingTimeToComplete_min"`
	ChargePowerKW                      int    `json:"chargePower_kW"`
	ChargeRateKmph                     int    `json:"chargeRate_kmph"`
}

// ChargingSettings is the /status.chargingSettings api
type ChargingSettings struct {
	CarCapturedTimestamp      string
	MaxChargeCurrentAC        string // reduced, maximum
	AutoUnlockPlugWhenCharged string
	TargetSOCPercent          int `json:"targetSOC_pct"`
}

// PlugStatus is the /status.plugStatus api
type PlugStatus struct {
	CarCapturedTimestamp string
	PlugConnectionState  string // connected, disconnected
	PlugLockState        string
}

// ClimatisationStatus is the /status.climatisationStatus api
type ClimatisationStatus struct {
	CarCapturedTimestamp          string
	RemainingClimatisationTimeMin int    `json:"remainingClimatisationTime_min"`
	ClimatisationState            string // off
}

// ClimatisationSettings is the /status.climatisationSettings api
type ClimatisationSettings struct {
	CarCapturedTimestamp              string
	TargetTemperatureK                float64 `json:"targetTemperature_K"`
	TargetTemperatureC                float64 `json:"targetTemperature_C"`
	ClimatisationWithoutExternalPower bool
	ClimatisationAtUnlock             bool // ClimatizationAtUnlock?
	WindowHeatingEnabled              bool
	ZoneFrontLeftEnabled              bool
	ZoneFrontRightEnabled             bool
	ZoneRearLeftEnabled               bool
	ZoneRearRightEnabled              bool
}

// RangeStatus is the /status.rangeStatus api
type RangeStatus struct {
	CarCapturedTimestamp string
	CarType              string
	PrimaryEngine        struct {
		Type              string
		CurrentSOCPercent int `json:"currentSOC_pct"`
		RemainingRangeKm  int `json:"remainingRange_km"`
	}
	TotalRangeKm int `json:"totalRange_km"`
}

// Status implements the /status response
func (v *API) Status(vin string) (res Status, err error) {
	uri := fmt.Sprintf("https://mobileapi.apps.emea.vwapps.io/vehicles/%s/status", vin)

	req, err := request.New(http.MethodGet, uri, nil, map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + "v.identity.Token()",
	})

	if err == nil {
		err = v.DoJSON(req, &res)
	}

	return res, err
}

// Action implements vehicle actions
func (v *API) Action(vin, action, value string) error {
	uri := fmt.Sprintf("https://mobileapi.apps.emea.vwapps.io/vehicles/%s/%s/%s", vin, action, value)

	req, err := request.New(http.MethodPost, uri, nil, map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + "v.identity.Token()",
	})

	if err == nil {
		var res interface{}
		err = v.DoJSON(req, &res)
	}

	return err
}

// Any implements any api response
func (v *API) Any(uri, vin string) (interface{}, error) {
	if strings.Contains(uri, "%s") {
		uri = fmt.Sprintf(uri, vin)
	}

	req, err := request.New(http.MethodGet, uri, nil, map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + "v.identity.Token()",
	})

	var res interface{}
	if err == nil {
		err = v.DoJSON(req, &res)
	}

	return res, err
}
