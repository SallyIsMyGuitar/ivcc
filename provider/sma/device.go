package sma

import (
	"fmt"
	"time"

	"github.com/andig/evcc/util"
	"github.com/imdario/mergo"
	"gitlab.com/bboehmke/sunny"
)

// Device holds information for a Device and provides interface to get values
type Device struct {
	*sunny.Device

	log    *util.Logger
	mux    *util.Waiter
	values map[sunny.ValueID]interface{}
}

func (d *Device) updateValues() error {
	d.mux.Lock()
	defer d.mux.Unlock()

	values, err := d.Device.GetValues()
	if err == nil {
		err = mergo.Merge(&d.values, values, mergo.WithOverride)
		d.mux.Update()
	}
	return err
}

func (d *Device) Values() (map[sunny.ValueID]interface{}, error) {
	elapsed := d.mux.LockWithTimeout()
	defer d.mux.Unlock()

	if elapsed > 0 {
		return nil, fmt.Errorf("update timeout: %v", elapsed.Truncate(time.Second))
	}

	return d.values, nil
}

func AsFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case nil:
		return 0
	default:
		util.NewLogger("sma").WARN.Printf("unknown value type: %T", value)
		return 0
	}
}
