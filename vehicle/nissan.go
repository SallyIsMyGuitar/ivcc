package vehicle

import (
	"fmt"
	"strings"
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
	"github.com/andig/evcc/vehicle/kamereon"
	"github.com/andig/evcc/vehicle/nissan"
)

// Credits to
//   https://github.com/Tobiaswk/dartnissanconnect
//   https://github.com/mitchellrj/kamereon-python
//   https://gitlab.com/tobiaswkjeldsen/carwingsflutter

// OAuth base url
// 	 https://prod.eu.auth.kamereon.org/kauth/oauth2/a-ncb-prod/.well-known/openid-configuration

// Nissan is an api.Vehicle implementation for Nissan cars
type Nissan struct {
	*embed
	*kamereon.Provider
}

func init() {
	registry.Add("nissan", NewNissanFromConfig)
}

// NewNissanFromConfig creates a new vehicle
func NewNissanFromConfig(other map[string]interface{}) (api.Vehicle, error) {
	cc := struct {
		embed               `mapstructure:",squash"`
		User, Password, VIN string
		Cache               time.Duration
	}{
		Cache: interval,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	v := &Nissan{
		embed: &cc.embed,
	}

	log := util.NewLogger("nissan")
	identity := nissan.NewIdentity(log)

	err := identity.Login(cc.User, cc.Password)
	if err != nil {
		return v, fmt.Errorf("login failed: %w", err)
	}

	api := nissan.NewAPI(log, identity, strings.ToUpper(cc.VIN))

	if err == nil && cc.VIN == "" {
		api.VIN, err = findVehicle(api.Vehicles())
		if err == nil {
			log.DEBUG.Printf("found vehicle: %v", api.VIN)
		}
	}

	v.Provider = kamereon.NewProvider(api.Battery, cc.Cache)

	return v, err
}
