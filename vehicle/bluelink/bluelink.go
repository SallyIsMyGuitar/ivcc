package bluelink

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/andig/evcc/provider"
	"github.com/andig/evcc/util"
	"github.com/google/uuid"
	"golang.org/x/net/publicsuffix"
)

const resOK = "S"

type Config struct {
	URI         string
	DeviceID    string
	Cookies     string
	Lang        string
	Login       string
	AccessToken string
	Vehicles    string
	SendPIN     string
	GetStatus   string
}

// API is an api.Vehicle implementation with configurable getters and setters.
type API struct {
	*util.HTTPHelper
	user     string
	password string
	pin      string
	chargeG  func() (float64, error)
	config   Config
	auth     Auth
}

type Auth struct {
	deviceID     string
	vehicleID    string
	controlToken string
}

type response struct {
	RetCode string `json:"retCode"`
	ResMsg  struct {
		DeviceID string `json:"deviceId"`
		EvStatus struct {
			BatteryStatus float64 `json:"batteryStatus"`
		} `json:"evStatus"`
		Vehicles []struct {
			VehicleID string `json:"vehicleId"`
		} `json:"vehicles"`
	} `json:"resMsg"`
}

// New creates a new BlueLink API
func New(log *util.Logger,
	user, password, pin string, cache time.Duration,
	config Config,
) (*API, error) {
	v := &API{
		HTTPHelper: util.NewHTTPHelper(log),
		config:     config,
		user:       user,
		password:   password,
		pin:        pin,
	}

	// api is unbelievably slowwhen retrieving status
	v.HTTPHelper.Client.Timeout = 60 * time.Second

	v.chargeG = provider.NewCached(v.chargeState, cache).FloatGetter()

	return v, nil
}

// request builds an HTTP request with headers and body
func (v *API) request(method, uri string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return req, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

// jsonRequest builds an HTTP json request with headers and body
func (v *API) jsonRequest(method, uri string, headers map[string]string, data interface{}) (*http.Request, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return v.request(method, uri, headers, bytes.NewReader(body))
}

// Credits to https://openwb.de/forum/viewtopic.php?f=5&t=1215&start=10#p11877

func (v *API) getDeviceID() (string, error) {
	uniID, _ := uuid.NewUUID()
	data := map[string]interface{}{
		"pushRegId": "1",
		"pushType":  "GCM",
		"uuid":      uniID.String(),
	}

	headers := map[string]string{
		"ccsp-service-id": "fdc85c00-0a2f-4c64-bcb4-2cfb1500730a",
		"Content-type":    "application/json;charset=UTF-8",
		"User-Agent":      "okhttp/3.10.0",
	}

	var resp response
	req, err := v.jsonRequest(http.MethodPost, v.config.URI+v.config.DeviceID, headers, data)
	if err == nil {
		_, err = v.RequestJSON(req, &resp)
	}

	return resp.ResMsg.DeviceID, err
}

func (v *API) getCookies() (cookieClient *util.HTTPHelper, err error) {
	cookieClient = util.NewHTTPHelper(v.Log)
	cookieClient.Client.Jar, err = cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})

	if err == nil {
		_, err = cookieClient.Get(v.config.URI + v.config.Cookies)
	}

	return cookieClient, err
}

func (v *API) setLanguage(cookieClient *util.HTTPHelper) error {
	headers := map[string]string{
		"Content-type": "application/json",
	}

	data := map[string]interface{}{
		"lang": "en",
	}

	req, err := v.jsonRequest(http.MethodPost, v.config.URI+v.config.Lang, headers, data)
	if err == nil {
		_, err = cookieClient.Request(req)
	}

	return err
}

func (v *API) login(cookieClient *util.HTTPHelper) (string, error) {
	headers := map[string]string{
		"Content-type": "application/json",
	}

	data := map[string]interface{}{
		"email":    v.user,
		"password": v.password,
	}

	req, err := v.jsonRequest(http.MethodPost, v.config.URI+v.config.Login, headers, data)
	if err != nil {
		return "", err
	}

	var redirect struct {
		RedirectURL string `json:"redirectUrl"`
	}

	var accCode string
	if _, err = cookieClient.RequestJSON(req, &redirect); err == nil {
		if parsed, err := url.Parse(redirect.RedirectURL); err == nil {
			accCode = parsed.Query().Get("code")
		}
	}

	return accCode, err
}

func (v *API) getToken(accCode string) (string, error) {
	headers := map[string]string{
		"Authorization": "Basic ZmRjODVjMDAtMGEyZi00YzY0LWJjYjQtMmNmYjE1MDA3MzBhOnNlY3JldA==",
		"Content-type":  "application/x-www-form-urlencoded",
		"User-Agent":    "okhttp/3.10.0",
	}

	data := "grant_type=authorization_code" +
		"&redirect_uri=https%3A%2F%2Fprd.eu-ccapi.kia.com%3A8080%2Fapi%2Fv1%2Fuser%2Foauth2%2Fredirect" +
		"&code=" + accCode

	req, err := v.request(http.MethodPost, v.config.URI+v.config.AccessToken, headers, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	var tokens struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	}

	var accToken string
	if _, err = v.RequestJSON(req, &tokens); err == nil {
		accToken = fmt.Sprintf("%s %s", tokens.TokenType, tokens.AccessToken)
	}

	return accToken, err
}

func (v *API) getVehicles(accToken, did string) (string, error) {
	headers := map[string]string{
		"Authorization":       accToken,
		"ccsp-device-id":      did,
		"ccsp-application-id": "693a33fa-c117-43f2-ae3b-61a02d24f417",
		"offset":              "1",
		"User-Agent":          "okhttp/3.10.0",
	}

	req, err := v.request(http.MethodGet, v.config.URI+v.config.Vehicles, headers, nil)
	if err == nil {
		var resp response
		if _, err = v.RequestJSON(req, &resp); err == nil {
			if len(resp.ResMsg.Vehicles) == 1 {
				return resp.ResMsg.Vehicles[0].VehicleID, nil
			}

			err = errors.New("couldn't find vehicle")
		}
	}

	return "", err
}

func (v *API) preWakeup(accToken, did, vid string) error {
	data := map[string]interface{}{
		"action":   "prewakeup",
		"deviceId": did,
	}

	headers := map[string]string{
		"Authorization":       accToken,
		"ccsp-device-id":      did,
		"ccsp-application-id": "693a33fa-c117-43f2-ae3b-61a02d24f417",
		"offset":              "1",
		"Content-Type":        "application/json;charset=UTF-8",
		"User-Agent":          "okhttp/3.10.0",
	}

	uri := v.config.URI + v.config.Vehicles + "/" + vid + "/control/engine"
	req, err := v.jsonRequest(http.MethodPost, uri, headers, data)
	if err == nil {
		_, err = v.Request(req)
	}

	return err
}

func (v *API) sendPIN(deviceID, accToken string) (string, error) {
	data := map[string]interface{}{
		"deviceId": deviceID,
		"pin":      string(v.pin),
	}

	headers := map[string]string{
		"Authorization": accToken,
		"Content-type":  "application/json;charset=UTF-8",
		"User-Agent":    "okhttp/3.10.0",
	}

	var token struct {
		ControlToken string `json:"controlToken"`
	}

	req, err := v.jsonRequest(http.MethodPut, v.config.URI+v.config.SendPIN, headers, data)
	if err == nil {
		_, err = v.RequestJSON(req, &token)
	}

	controlToken := ""
	if err == nil {
		controlToken = "Bearer " + token.ControlToken

	}

	return controlToken, err
}

func (v *API) authFlow() (err error) {
	v.auth.deviceID, err = v.getDeviceID()

	var cookieClient *util.HTTPHelper
	if err == nil {
		cookieClient, err = v.getCookies()
	}

	if err == nil {
		err = v.setLanguage(cookieClient)
	}

	var kiaAccCode string
	if err == nil {
		kiaAccCode, err = v.login(cookieClient)
	}

	var kiaAccToken string
	if err == nil {
		kiaAccToken, err = v.getToken(kiaAccCode)
	}

	if err == nil {
		v.auth.vehicleID, err = v.getVehicles(kiaAccToken, v.auth.deviceID)
	}

	if err == nil {
		err = v.preWakeup(kiaAccToken, v.auth.deviceID, v.auth.vehicleID)
	}

	if err == nil {
		v.auth.controlToken, err = v.sendPIN(v.auth.deviceID, kiaAccToken)
	}

	return err
}

func (v *API) getStatus() (float64, error) {
	headers := map[string]string{
		"Authorization":  v.auth.controlToken,
		"ccsp-device-id": v.auth.deviceID,
		"Content-Type":   "application/json",
	}

	var resp response
	uri := v.config.URI + v.config.GetStatus + "/" + v.auth.vehicleID + "/status"
	req, err := v.request(http.MethodGet, uri, headers, nil)
	if err == nil {
		_, err = v.RequestJSON(req, &resp)

		if err == nil && resp.RetCode != resOK {
			err = errors.New("unexpected response")
		}
	}

	return resp.ResMsg.EvStatus.BatteryStatus, err
}

// chargeState implements the Vehicle.ChargeState interface
func (v *API) chargeState() (float64, error) {
	soc, err := v.getStatus()

	if err != nil && v.HTTPHelper.LastResponse().StatusCode == http.StatusUnauthorized {
		if err = v.authFlow(); err == nil {
			soc, err = v.getStatus()
		}
	}

	return soc, err
}

// ChargeState implements the Vehicle.ChargeState interface
func (v *API) ChargeState() (float64, error) {
	return v.chargeG()
}
