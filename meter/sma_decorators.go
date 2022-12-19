package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateSMA(base *SMA, battery func() (float64, error)) api.Meter {
	switch {
	case battery == nil:
		return base

	case battery != nil:
		return &struct {
			*SMA
			api.Battery
		}{
			SMA: base,
			Battery: &decorateSMABatteryImpl{
				battery: battery,
			},
		}
	}

	return nil
}

type decorateSMABatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateSMABatteryImpl) Soc() (float64, error) {
	return impl.battery()
}
