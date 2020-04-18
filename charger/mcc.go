package charger

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/andig/evcc/api"
)

const (
	apiLogin                   apiFunction = "jwt/login"
	apiRefresh                 apiFunction = "jwt/refresh"
	apiChargeState             apiFunction = "v1/api/WebServer/properties/chargeState"
	apiCurrentSession          apiFunction = "v1/api/WebServer/properties/swaggerCurrentSession"
	apiEnergy                  apiFunction = "v1/api/iCAN/properties/propjIcanEnergy"
	apiSetCurrentLimit         apiFunction = "v1/api/SCC/properties/propHMICurrentLimit?value="
	apiCurrentCableInformation apiFunction = "v1/api/SCC/properties/json_CurrentCableInformation"
)

// MCCErrorResponse is the API response if status not OK
type MCCErrorResponse struct {
	Error string
}

// MCCTokenResponse is the apiLogin response
type MCCTokenResponse struct {
	Token string
}

// MCCCurrentSession is the apiCurrentSession response
type MCCCurrentSession struct {
	Duration     time.Duration
	EnergySumKwh float64
}

// MCCEnergyPhase is the apiEnergy response for a single phase
type MCCEnergyPhase struct {
	Power float64
}

// MCCEnergy is the apiEnergy response
type MCCEnergy struct {
	L1, L2, L3 MCCEnergyPhase
}

// MCCCurrentCableInformation is the apiCurrentCableInformation response
type MCCCurrentCableInformation struct {
	MaxValue, MinValue, Value int64
}

// MCC charger implementation for supporting Mobile Charger Connect devices from Audi, Bentley, Porsche
type MobileConnect struct {
	*api.HTTPHelper
	uri              string
	password         string
	token            string
	tokenValid       time.Time
	tokenRefresh     time.Time
	cableInformation MCCCurrentCableInformation
}

// NewMobileConnectFromConfig creates a MCC charger from generic config
func NewMobileConnectFromConfig(log *api.Logger, other map[string]interface{}) api.Charger {
	cc := struct{ URI, Password string }{}
	api.DecodeOther(log, other, &cc)

	return NewMobileConnect(cc.URI, cc.Password)
}

// NewMobileConnect creates MCC charger
func NewMobileConnect(uri string, password string) *MobileConnect {
	// ignore the self signed certificate
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Transport: customTransport}

	mcc := &MobileConnect{
		HTTPHelper: api.NewHTTPHelperWithClient(api.NewLogger("mcc "), client),
		uri:        strings.TrimRight(uri, "/"),
		password:   password,
	}

	return mcc
}

// construct the URL for a given apiFunction
func (mcc *MobileConnect) apiURL(api apiFunction) string {
	return fmt.Sprintf("%s/%s", mcc.uri, api)
}

// proces the http request to fetch the auth token for a login or refresh request
func (mcc *MobileConnect) fetchToken(request *http.Request) error {
	var tr MCCTokenResponse
	b, err := mcc.RequestJSON(request, &tr)
	if err == nil {
		if len(tr.Token) == 0 && len(b) > 0 {
			var error MCCErrorResponse

			if err := json.Unmarshal(b, &error); err != nil {
				return err
			}

			return fmt.Errorf("response: %s", error.Error)
		}

		mcc.token = tr.Token
		// According to the Web Interface, the token is valid for 10 minutes
		mcc.tokenValid = time.Now().Add(10 * time.Minute)

		// the web interface updates the token every 2 minutes, so lets do the same here
		mcc.tokenRefresh = time.Now().Add(2 * time.Minute)
	}

	return err
}

// login as the home user with the given password
func (mcc *MobileConnect) login(password string) error {
	uri := fmt.Sprintf("%s/%s", mcc.uri, apiLogin)

	data := url.Values{
		"user": []string{"user"},
		"pass": []string{mcc.password},
	}

	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return mcc.fetchToken(req)
}

// refresh the auth token with a new one
func (mcc *MobileConnect) refresh() error {
	uri := fmt.Sprintf("%s/%s", mcc.uri, apiRefresh)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mcc.token))

	return mcc.fetchToken(req)
}

// creates a http request that contains the auth token
func (mcc *MobileConnect) request(method, uri string) (*http.Request, error) {
	// do we need to login?
	if mcc.token == "" || time.Since(mcc.tokenValid) > 0 {
		if err := mcc.login(mcc.password); err != nil {
			return nil, err
		}
	}

	// is it time to refresh the token?
	if time.Since(mcc.tokenRefresh) > 0 {
		if err := mcc.refresh(); err != nil {
			return nil, err
		}
	}

	// now lets process the request with the fetched token
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return req, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", mcc.token))

	return req, nil
}

// use http PUT to set a value on the URI, the value should be URL encoded in the URI parameter
func (mcc *MobileConnect) putValue(uri string) ([]byte, error) {
	req, err := mcc.request(http.MethodPut, uri)
	if err != nil {
		return nil, err
	}

	return mcc.Request(req)
}

// use http GET to fetch a non structured value from an URI and stores it in result
func (mcc *MobileConnect) getValue(uri string) ([]byte, error) {
	req, err := mcc.request(http.MethodGet, uri)
	if err != nil {
		return nil, err
	}

	return mcc.Request(req)
}

// use http GET to fetch an escaped JSON string and unmarshall the data in result
func (mcc *MobileConnect) getEscapedJSON(uri string, result interface{}) error {
	req, err := mcc.request(http.MethodGet, uri)
	if err != nil {
		return err
	}

	b, err := mcc.Request(req)
	if err == nil {
		s, _ := strconv.Unquote(strings.Trim(string(b), "\n"))
		if err := json.Unmarshal([]byte(s), &result); err != nil {
			return err
		}
	}

	return err
}

// Status implements the Charger.Status interface
func (mcc *MobileConnect) Status() (api.ChargeStatus, error) {
	b, err := mcc.getValue(mcc.apiURL(apiChargeState))
	if err != nil {
		return api.StatusNone, err
	}

	chargeState, err := strconv.ParseInt(strings.Trim(string(b), "\n"), 10, 8)
	if err != nil {
		return api.StatusNone, err
	}

	switch chargeState {
	case 0:
		// 0: Unplugged		StatusA
		return api.StatusA, nil
	case 1:
		// 1: Connecting	StatusB
		return api.StatusB, nil
	case 2:
		// 2: Error				StatusE
		return api.StatusE, nil
	case 3:
		// 3: Established	StatusF
		return api.StatusF, nil
	case 4, 6:
		// 4: Paused			StatusD
		// 6: Finished		StatusD
		return api.StatusD, nil
	case 5:
		// 5: Active			StatusC
		return api.StatusC, nil
	default:
		return api.StatusNone, fmt.Errorf("properties unknown result: %d", chargeState)
	}
}

// Enabled implements the Charger.Enabled interface
func (mcc *MobileConnect) Enabled() (bool, error) {
	// Check if the car is connected and Paused, Active, or Finished
	b, err := mcc.getValue(mcc.apiURL(apiChargeState))
	if err != nil {
		return false, err
	}

	// return value is returned in the format 0\n
	chargeState, err := strconv.ParseInt(strings.Trim(string(b), "\n"), 10, 8)
	if err != nil {
		return false, err
	}

	if chargeState >= 4 && chargeState <= 6 {
		return true, nil
	}

	return false, nil
}

// Enable implements the Charger.Enable interface
func (mcc *MobileConnect) Enable(enable bool) error {
	// As we don't know of the API to disable charging this for now always returns an error
	return nil
}

// MaxCurrent implements the Charger.MaxCurrent interface
func (mcc *MobileConnect) MaxCurrent(current int64) error {
	// The device doesn't return an error if we set a value greater than the
	// current allowed max or smaller than the allowed min
	// instead it will simply set it to max or min and return "OK" anyway
	// Since the API here works differently, we fetch the limits
	// and then return an error if the value is outside of the limits or
	// otherwise set the new value
	if mcc.cableInformation.MaxValue == 0 {
		if err := mcc.getEscapedJSON(mcc.apiURL(apiCurrentCableInformation), &mcc.cableInformation); err != nil {
			return err
		}
	}

	if current < mcc.cableInformation.MinValue {
		return fmt.Errorf("value is lower than the allowed minimum value %d", mcc.cableInformation.MinValue)
	}

	if current > mcc.cableInformation.MaxValue {
		return fmt.Errorf("value is higher than the allowed maximum value %d", mcc.cableInformation.MaxValue)
	}

	url := fmt.Sprintf("%s%d", mcc.apiURL(apiSetCurrentLimit), current)

	b, err := mcc.putValue(url)
	if err != nil {
		return err
	}

	// return value is returned in the format "OK"\n
	s := strings.Trim(string(b), "\n")
	s = strings.Trim(s, "\"")

	if s != "OK" {
		return fmt.Errorf("Call returned an unexpected error")
	}

	return nil
}

// CurrentPower implements the Meter interface.
func (mcc *MobileConnect) CurrentPower() (float64, error) {
	var energy MCCEnergy
	err := mcc.getEscapedJSON(mcc.apiURL(apiEnergy), &energy)

	return energy.L1.Power + energy.L2.Power + energy.L3.Power, err
}

// ChargedEnergy implements the ChargeRater interface.
func (mcc *MobileConnect) ChargedEnergy() (float64, error) {
	var currentSession MCCCurrentSession
	if err := mcc.getEscapedJSON(mcc.apiURL(apiCurrentSession), &currentSession); err != nil {
		return 0, err
	}

	return currentSession.EnergySumKwh, nil
}

// ChargingTime yields current charge run duration
func (mcc *MobileConnect) ChargingTime() (time.Duration, error) {
	var currentSession MCCCurrentSession
	if err := mcc.getEscapedJSON(mcc.apiURL(apiCurrentSession), &currentSession); err != nil {
		return 0, err
	}

	return time.Duration(currentSession.Duration * time.Second), nil
}
