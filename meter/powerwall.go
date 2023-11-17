package meter

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/andig/go-powerwall"
	"github.com/bogosj/tesla"
	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/provider"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"golang.org/x/oauth2"
)

// PowerWall is the tesla powerwall meter
type PowerWall struct {
	usage                 string
	client                *powerwall.Client
	meterG                func() (map[string]powerwall.MeterAggregatesData, error)
	energySite            *tesla.EnergySite
	batteryControl        bool
	defaultBatteryReserve uint
}

func init() {
	registry.Add("tesla", NewPowerWallFromConfig)
	registry.Add("powerwall", NewPowerWallFromConfig)
}

//go:generate go run ../cmd/tools/decorate.go -f decoratePowerWall -b *PowerWall -r api.Meter -t "api.MeterEnergy,TotalEnergy,func() (float64, error)" -t "api.Battery,Soc,func() (float64, error)" -t "api.BatteryCapacity,Capacity,func() float64"

// NewPowerWallFromConfig creates a PowerWall Powerwall Meter from generic config
func NewPowerWallFromConfig(other map[string]interface{}) (api.Meter, error) {
	cc := struct {
		URI, Usage, User, Password string
		Cache                      time.Duration
		RefreshToken               string
		EnergySiteProdId           int64
		DefaultBatteryReserve      uint
	}{
		Cache:                 time.Second,
		DefaultBatteryReserve: 20,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Usage == "" {
		return nil, errors.New("missing usage")
	}

	if cc.Password == "" {
		return nil, errors.New("missing password")
	}

	var batteryControl bool
	if cc.RefreshToken != "" || cc.EnergySiteProdId != 0 {
		if cc.RefreshToken == "" {
			return nil, errors.New("missing refresh token")
		}
		batteryControl = true
	}

	// support default meter names
	switch strings.ToLower(cc.Usage) {
	case "grid":
		cc.Usage = "site"
	case "pv":
		cc.Usage = "solar"
	}

	return NewPowerWall(cc.URI, cc.Usage, cc.User, cc.Password, cc.Cache, cc.RefreshToken, cc.EnergySiteProdId, cc.DefaultBatteryReserve, batteryControl)
}

// NewPowerWall creates a Tesla PowerWall Meter
func NewPowerWall(uri, usage, user, password string, cache time.Duration, refreshToken string, energySiteProdId int64, defaultBatteryReserve uint, batteryControl bool) (api.Meter, error) {
	log := util.NewLogger("powerwall").Redact(user, password, refreshToken)

	httpClient := &http.Client{
		Transport: request.NewTripper(log, powerwall.DefaultTransport()),
		Timeout:   time.Second * 2, // Timeout after 2 seconds
	}

	client := powerwall.NewClient(uri, user, password, powerwall.WithHttpClient(httpClient))
	if _, err := client.GetStatus(); err != nil {
		return nil, err
	}

	m := &PowerWall{
		client:         client,
		usage:          strings.ToLower(usage),
		meterG:         provider.Cached(client.GetMetersAggregates, cache),
		batteryControl: batteryControl,
	}

	if batteryControl {
		ctx := context.WithValue(context.Background(), oauth2.HTTPClient, request.NewClient(log))

		options := []tesla.ClientOption{tesla.WithToken(&oauth2.Token{
			RefreshToken: refreshToken,
			Expiry:       time.Now(),
		})}

		cloudClient, err := tesla.NewClient(ctx, options...)
		if err != nil {
			return nil, err
		}

		if energySiteProdId == 0 {
			//auto detect energy site ID, picking first
			products, err := cloudClient.Products()
			if err != nil {
				return nil, err
			}

			for _, p := range products {
				if p.EnergySiteId != 0 {
					energySiteProdId = p.EnergySiteId
					break
				}
			}
			log.INFO.Printf("Auto-detected Energy Site with Product ID %d", energySiteProdId)
		}

		energySite, err := cloudClient.EnergySite(energySiteProdId)
		if err != nil {
			return nil, err
		}
		m.energySite = energySite

		m.defaultBatteryReserve = defaultBatteryReserve
	}

	// decorate api.MeterEnergy
	var totalEnergy func() (float64, error)
	if m.usage == "load" || m.usage == "solar" {
		totalEnergy = m.totalEnergy
	}

	// decorate api.BatterySoc
	var batterySoc func() (float64, error)
	var batteryCapacity func() float64
	if usage == "battery" {
		batterySoc = m.batterySoc

		res, err := m.client.GetSystemStatus()
		if err != nil {
			return nil, err
		}

		batteryCapacity = func() float64 {
			return float64(res.NominalFullPackEnergy) / 1e3
		}
	}

	return decoratePowerWall(m, totalEnergy, batterySoc, batteryCapacity), nil
}

var _ api.Meter = (*PowerWall)(nil)

// CurrentPower implements the api.Meter interface
func (m *PowerWall) CurrentPower() (float64, error) {
	res, err := m.meterG()
	if err != nil {
		return 0, err
	}

	if o, ok := res[m.usage]; ok {
		return float64(o.InstantPower), nil
	}

	return 0, fmt.Errorf("invalid usage: %s", m.usage)
}

// totalEnergy implements the api.MeterEnergy interface
func (m *PowerWall) totalEnergy() (float64, error) {
	res, err := m.meterG()
	if err != nil {
		return 0, err
	}

	if o, ok := res[m.usage]; ok {
		switch m.usage {
		case "load":
			return float64(o.EnergyImported) / 1e3, nil
		case "solar":
			return float64(o.EnergyExported) / 1e3, nil
		}
	}

	return 0, fmt.Errorf("invalid usage: %s", m.usage)
}

// batterySoc implements the api.Battery interface
func (m *PowerWall) batterySoc() (float64, error) {
	res, err := m.client.GetSOE()
	if err != nil {
		return 0, err
	}

	return float64(res.Percentage), err
}

// SetBatteryMode implements the api.BatteryController interface
func (m *PowerWall) SetBatteryMode(mode api.BatteryMode) error {
	if !m.batteryControl {
		return nil
	}

	switch mode {
	case api.BatteryNormal: // set minSoc to cofigured default
		if err := m.energySite.SetBatteryReserve(uint64(m.defaultBatteryReserve)); err != nil {
			return err
		}
	case api.BatteryHold: // set minSoc to currentSoc
		ess, err := m.energySite.EnergySiteStatus()
		if err != nil {
			return err
		}
		currentSoc := math.Round(ess.PercentageCharged + 0.5) // .5 ensures we round up
		if err := m.energySite.SetBatteryReserve(uint64(currentSoc)); err != nil {
			return err
		}
	case api.BatteryCharge: // set mincSoc to 100; will not cause grid charge, but any PV surplus will charge
		if err := m.energySite.SetBatteryReserve(100); err != nil {
			return err
		}
	}
	return nil
}
