package provider

import (
	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
)

type openWBStatusProvider struct {
	plugged, charging BoolGetter
}

func openWBStatusFromConfig(log *util.Logger, other map[string]interface{}) StringGetter {
	cc := struct {
		Plugged, Charging Config
	}{}
	util.DecodeOther(log, other, &cc)

	o := &openWBStatusProvider{
		plugged:  NewBoolGetterFromConfig(log, cc.Plugged),
		charging: NewBoolGetterFromConfig(log, cc.Charging),
	}

	return o.stringGetter
}

func (o *openWBStatusProvider) stringGetter() (string, error) {
	charging, err := o.charging()
	if err != nil {
		return "", err
	}
	if charging {
		return string(api.StatusC), nil
	}

	plugged, err := o.plugged()
	if err != nil {
		return "", err
	}
	if plugged {
		return string(api.StatusB), nil
	}

	return string(api.StatusA), nil
}
