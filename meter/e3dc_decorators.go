package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateE3dc(base *E3dc, battery func() (float64, error)) api.Meter {
	switch {
	case battery == nil:
		return base

	case battery != nil:
		return &struct {
			*E3dc
			api.Battery
		}{
			E3dc: base,
			Battery: &decorateE3dcBatteryImpl{
				battery: battery,
			},
		}
	}

	return nil
}

type decorateE3dcBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateE3dcBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}
