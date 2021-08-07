package charger

import (
	"encoding/binary"
	"fmt"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
	"github.com/andig/evcc/util/modbus"
)

// HeidelbergEC charger implementation
type HeidelbergEC struct {
	log     *util.Logger
	conn    *modbus.Connection
	current uint16
}

const (
	hecRegFirmware      = 1   // Input
	hecRegVehicleStatus = 5   // Input
	hecRegPower         = 14  // Input
	hecRegEnergy        = 17  // Input
	hecRegAmpsConfig    = 261 // Holding
)

var hecRegCurrents = []uint16{6, 7, 8}

func init() {
	registry.Add("heidelberg", NewHeidelbergECFromConfig)
}

// https://wallbox.heidelberg.com/wp-content/uploads/2021/05/EC_ModBus_register_table_20210222.pdf (newer)
// https://cdn.shopify.com/s/files/1/0101/2409/9669/files/heidelberg-energy-control-modbus.pdf (older)

// NewHeidelbergECFromConfig creates a HeidelbergEC charger from generic config
func NewHeidelbergECFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		modbus.Settings `mapstructure:",squash"`
	}{
		Settings: modbus.Settings{
			ID: 1,
		},
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return NewHeidelbergEC(cc.URI, cc.Device, cc.Comset, cc.Baudrate, true, cc.ID)
}

// NewHeidelbergEC creates HeidelbergEC charger
func NewHeidelbergEC(uri, device, comset string, baudrate int, rtu bool, slaveID uint8) (api.Charger, error) {
	conn, err := modbus.NewConnection(uri, device, comset, baudrate, rtu, slaveID)
	if err != nil {
		return nil, err
	}

	log := util.NewLogger("hec")
	conn.Logger(log.TRACE)

	wb := &HeidelbergEC{
		log:     log,
		conn:    conn,
		current: 60, // assume min current
	}

	return wb, nil
}

// Status implements the api.Charger interface
func (wb *HeidelbergEC) Status() (api.ChargeStatus, error) {
	b, err := wb.conn.ReadInputRegisters(hecRegVehicleStatus, 1)
	if err != nil {
		return api.StatusNone, err
	}

	switch b[0] {
	case 2, 3:
		return api.StatusA, nil
	case 4, 5:
		return api.StatusB, nil
	case 6, 7:
		return api.StatusC, nil
	default:
		return api.StatusNone, fmt.Errorf("invalid status: %v", b)
	}
}

// Enabled implements the api.Charger interface
func (wb *HeidelbergEC) Enabled() (bool, error) {
	b, err := wb.conn.ReadHoldingRegisters(hecRegAmpsConfig, 1)
	if err != nil {
		return false, err
	}

	cur := binary.BigEndian.Uint16(b)

	enabled := cur != 0
	if enabled {
		wb.current = cur
	}

	return enabled, nil
}

// Enable implements the api.Charger interface
func (wb *HeidelbergEC) Enable(enable bool) error {
	var cur uint16
	if enable {
		cur = wb.current
	}

	_, err := wb.conn.WriteSingleRegister(hecRegAmpsConfig, cur)

	return err
}

// MaxCurrent implements the api.Charger interface
func (wb *HeidelbergEC) MaxCurrent(current int64) error {
	if current < 6 {
		return fmt.Errorf("invalid current %d", current)
	}

	return wb.MaxCurrentMillis(float64(current))
}

var _ api.ChargerEx = (*HeidelbergEC)(nil)

// MaxCurrentMillis implements the api.ChargerEx interface
func (wb *HeidelbergEC) MaxCurrentMillis(current float64) error {
	if current < 6 {
		return fmt.Errorf("invalid current %.1f", current)
	}

	cur := uint16(10 * current)

	_, err := wb.conn.WriteSingleRegister(hecRegAmpsConfig, cur)
	if err == nil {
		wb.current = cur
	}

	return err
}

var _ api.Meter = (*HeidelbergEC)(nil)

// CurrentPower implements the api.Meter interface
func (wb *HeidelbergEC) CurrentPower() (float64, error) {
	b, err := wb.conn.ReadInputRegisters(hecRegPower, 1)
	if err != nil {
		return 0, err
	}

	return float64(binary.BigEndian.Uint16(b)), nil
}

var _ api.MeterEnergy = (*HeidelbergEC)(nil)

// TotalEnergy implements the api.MeterEnergy interface
func (wb *HeidelbergEC) TotalEnergy() (float64, error) {
	b, err := wb.conn.ReadInputRegisters(hecRegEnergy, 2)
	if err != nil {
		return 0, err
	}

	return float64(binary.BigEndian.Uint32(b)), nil
}

var _ api.MeterCurrent = (*HeidelbergEC)(nil)

// Currents implements the api.MeterCurrent interface
func (wb *HeidelbergEC) Currents() (float64, float64, float64, error) {
	var currents []float64
	for _, regCurrent := range hecRegCurrents {
		b, err := wb.conn.ReadInputRegisters(regCurrent, 2)
		if err != nil {
			return 0, 0, 0, err
		}

		currents = append(currents, float64(binary.BigEndian.Uint16(b))/10)
	}

	return currents[0], currents[1], currents[2], nil
}

var _ api.Diagnosis = (*HeidelbergEC)(nil)

// Diagnose implements the api.Diagnosis interface
func (wb *HeidelbergEC) Diagnose() {
	b, err := wb.conn.ReadInputRegisters(hecRegFirmware, 2)
	if err == nil {
		fmt.Printf("Firmware:\t%d.%d.%d\n", b[1], b[2], b[3])
	}
}
