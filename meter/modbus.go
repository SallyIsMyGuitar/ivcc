package meter

import (
	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
	"github.com/andig/evcc/util/modbus"
	"github.com/volkszaehler/mbmd/meters"
	"github.com/volkszaehler/mbmd/meters/rs485"
	"github.com/volkszaehler/mbmd/meters/sunspec"
)

// Modbus is an api.Meter implementation with configurable getters and setters.
type Modbus struct {
	log      *util.Logger
	conn     meters.Connection
	device   meters.Device
	slaveID  uint8
	opPower  modbus.Operation
	opEnergy modbus.Operation
}

// NewModbusFromConfig creates api.Meter from config
func NewModbusFromConfig(log *util.Logger, other map[string]interface{}) api.Meter {
	cc := struct {
		modbus.Settings `mapstructure:",squash"`
		Power, Energy   string
	}{}
	util.DecodeOther(log, other, &cc)

	// assume RTU if not set and this is a known RS485 meter model
	if cc.RTU == nil {
		b := modbus.IsRS485(cc.Model)
		cc.RTU = &b
	}

	conn := modbus.NewConnection(log, cc.URI, cc.Device, cc.Comset, cc.Baudrate, *cc.RTU)
	device, err := modbus.NewDevice(log, cc.Model, *cc.RTU)

	log = util.NewLogger("modb")
	conn.Logger(log.TRACE)

	// prepare device
	if err == nil {
		conn.Slave(cc.ID)
		err = device.Initialize(conn.ModbusClient())

		// silence Kostal implementation errors
		if _, partial := err.(meters.SunSpecPartiallyInitialized); partial {
			err = nil
		}
	}
	if err != nil {
		log.FATAL.Fatal(err)
	}

	m := &Modbus{
		log:     log,
		conn:    conn,
		device:  device,
		slaveID: cc.ID,
	}

	// power reading
	if cc.Power == "" {
		cc.Power = "Power"
	}

	if err := modbus.ParseOperation(device, cc.Power, &m.opPower); err != nil {
		log.FATAL.Fatalf("invalid measurement for power: %s", cc.Power)
	}

	// decorate energy reading
	if cc.Energy != "" {
		if err := modbus.ParseOperation(device, cc.Energy, &m.opEnergy); err != nil {
			log.FATAL.Fatalf("invalid measurement for energy: %s", cc.Power)
		}

		return &ModbusEnergy{m}
	}

	return m
}

// floatGetter executes configured modbus read operation and implements func() (float64, error)
func (m *Modbus) floatGetter(op modbus.Operation) (float64, error) {
	m.conn.Slave(m.slaveID)

	var res meters.MeasurementResult
	var err error

	if dev, ok := m.device.(*rs485.RS485); ok {
		res, err = dev.QueryOp(m.conn.ModbusClient(), op.MBMD)
	}

	if dev, ok := m.device.(*sunspec.SunSpec); ok {
		if op.MBMD.IEC61850 != 0 {
			res, err = dev.QueryOp(m.conn.ModbusClient(), op.MBMD.IEC61850)
		} else {
			res, err = dev.QueryPoint(
				m.conn.ModbusClient(),
				op.SunSpec.Model,
				op.SunSpec.Block,
				op.SunSpec.Point,
			)
		}
	}

	if err == nil {
		m.log.TRACE.Printf("%+v", res)
	}

	return res.Value, err
}

// CurrentPower implements the Meter.CurrentPower interface
func (m *Modbus) CurrentPower() (float64, error) {
	return m.floatGetter(m.opPower)
}

// ModbusEnergy decorates Modbus with api.MeterEnergy interface
type ModbusEnergy struct {
	*Modbus
}

// TotalEnergy implements the Meter.TotalEnergy interface
func (m *ModbusEnergy) TotalEnergy() (float64, error) {
	return m.floatGetter(m.opEnergy)
}
