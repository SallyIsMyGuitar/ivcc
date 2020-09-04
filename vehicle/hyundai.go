package vehicle

import (
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
	"github.com/andig/evcc/vehicle/bluelink"
)

// Hyundai is an api.Vehicle implementation
type Hyundai struct {
	*embed
	*bluelink.API
}

func init() {
	registry.Add("hyundai", NewHyundaiFromConfig)
}

// NewHyundaiFromConfig creates a new Vehicle
func NewHyundaiFromConfig(other map[string]interface{}) (api.Vehicle, error) {
	cc := struct {
		Title          string
		Capacity       int64
		User, Password string
		PIN            string
		Cache          time.Duration
	}{}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	settings := bluelink.Config{
		URI:               "https://prd.eu-ccapi.hyundai.com:8080",
		TokenAuth:         "NmQ0NzdjMzgtM2NhNC00Y2YzLTk1NTctMmExOTI5YTk0NjU0OktVeTQ5WHhQekxwTHVvSzB4aEJDNzdXNlZYaG10UVI5aVFobUlGampvWTRJcHhzVg==",
		CCSPServiceID:     "6d477c38-3ca4-4cf3-9557-2a1929a94654",
		CCSPApplicationID: "99cfff84-f4e2-4be8-a5ed-e5b755eb6581",
		DeviceID:          "/api/v1/spa/notifications/register",
		Lang:              "/api/v1/user/language",
		Login:             "/api/v1/user/signin",
		AccessToken:       "/api/v1/user/oauth2/token",
		Vehicles:          "/api/v1/spa/vehicles",
		SendPIN:           "/api/v1/user/pin",
		GetStatus:         "/api/v2/spa/vehicles/",
	}

	log := util.NewLogger("hyundai")
	api, err := bluelink.New(log, cc.User, cc.Password, cc.PIN, cc.Cache, settings)
	if err != nil {
		return nil, err
	}

	v := &Hyundai{
		embed: &embed{cc.Title, cc.Capacity},
		API:   api,
	}

	return v, nil
}
