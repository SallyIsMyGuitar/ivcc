package meter

import (
	"errors"
	"fmt"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/internal"
	"github.com/andig/evcc/provider"
	"github.com/andig/evcc/util"
)

func init() {
	registry.Add("default", NewConfigurableFromConfig)
	registry.Add(internal.Custom, NewConfigurableFromConfig)
}

//go:generate go run ../../cmd/tools/decorate.go -f decorateMeter -b api.Meter -t "api.MeterEnergy,TotalEnergy,func() (float64, error)" -t "api.MeterCurrent,Currents,func() (float64, float64, float64, error)" -t "api.Battery,SoC,func() (float64, error)"

// NewConfigurableFromConfig creates api.Meter from config
func NewConfigurableFromConfig(other map[string]interface{}) (api.Meter, error) {
	cc := struct {
		Power    provider.Config
		Energy   *provider.Config  // optional
		SoC      *provider.Config  // optional
		Currents []provider.Config // optional
	}{}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	for k, v := range map[string]string{"power": cc.Power.PluginType()} {
		if v == "" {
			return nil, fmt.Errorf("missing plugin configuration: %s", k)
		}
	}

	power, err := provider.NewFloatGetterFromConfig(cc.Power)
	if err != nil {
		return nil, fmt.Errorf("power: %w", err)
	}

	m, _ := NewConfigurable(power)

	// decorate Meter with MeterEnergy
	if cc.Energy != nil {
		m.totalEnergyG, err = provider.NewFloatGetterFromConfig(*cc.Energy)
		if err != nil {
			return nil, fmt.Errorf("energy: %w", err)
		}
	}

	// decorate Meter with MeterCurrent
	if len(cc.Currents) > 0 {
		if len(cc.Currents) != 3 {
			return nil, errors.New("need 3 currents")
		}

		for idx, cc := range cc.Currents {
			c, err := provider.NewFloatGetterFromConfig(cc)
			if err != nil {
				return nil, fmt.Errorf("currents[%d]: %w", idx, err)
			}

			m.currentsG = append(m.currentsG, c)
		}
	}

	// decorate Meter with BatterySoC
	if cc.SoC != nil {
		m.batterySoCG, err = provider.NewFloatGetterFromConfig(*cc.SoC)
		if err != nil {
			return nil, fmt.Errorf("battery: %w", err)
		}
	}

	res := m.Decorate(m.totalEnergyG, m.currentsG, m.batterySoCG)

	return res, nil
}

// NewConfigurable creates a new meter
func NewConfigurable(currentPowerG func() (float64, error)) (*Meter, error) {
	m := &Meter{
		currentPowerG: currentPowerG,
	}
	return m, nil
}

// Meter is an api.Meter implementation with configurable getters and setters.
type Meter struct {
	currentPowerG func() (float64, error)
	totalEnergyG  func() (float64, error)
	currentsG     []func() (float64, error)
	batterySoCG   func() (float64, error)
}

// Decorate attaches additional capabilities to the base meter
func (m *Meter) Decorate(
	totalEnergyG func() (float64, error),
	currentsG []func() (float64, error),
	batterySoCG func() (float64, error),
) api.Meter {
	var totalEnergy func() (float64, error)
	if totalEnergyG != nil {
		m.totalEnergyG = totalEnergyG
		totalEnergy = m.totalEnergy
	}

	var currents func() (float64, float64, float64, error)
	if currentsG != nil {
		m.currentsG = currentsG
		currents = m.currents
	}

	var batterySoC func() (float64, error)
	if batterySoCG != nil {
		m.batterySoCG = batterySoCG
		batterySoC = m.batterySoC
	}

	return decorateMeter(m, totalEnergy, currents, batterySoC)
}

// CurrentPower implements the api.Meter interface
func (m *Meter) CurrentPower() (float64, error) {
	return m.currentPowerG()
}

// totalEnergy implements the api.MeterEnergy interface
func (m *Meter) totalEnergy() (float64, error) {
	return m.totalEnergyG()
}

// currents implements the api.MeterCurrent interface
func (m *Meter) currents() (float64, float64, float64, error) {
	var currents []float64
	for _, currentG := range m.currentsG {
		c, err := currentG()
		if err != nil {
			return 0, 0, 0, err
		}

		currents = append(currents, c)
	}

	return currents[0], currents[1], currents[2], nil
}

// batterySoC implements the api.Battery interface
func (m *Meter) batterySoC() (float64, error) {
	return m.batterySoCG()
}
