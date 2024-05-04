package modbus

import (
	"sync"
	"time"

	"github.com/grid-x/modbus"
	"github.com/volkszaehler/mbmd/meters"
)

// Connection is a logical modbus connection per slave ID sharing a physical connection
type Connection struct {
	meters.Connection
	mu      sync.Mutex
	logger  *logger
	logical meters.Logger
	delay   func()
}

func (c *Connection) Delay(delay time.Duration) {
	c.delay = func() { time.Sleep(delay) }
}

func (c *Connection) Clone(slaveID uint8) *Connection {
	return &Connection{
		Connection: c.Connection.Clone(slaveID),
	}
}

// TODO
func (c *Connection) ConnectDelay(_ time.Duration) {
}

// TODO
func (c *Connection) Timeout(_ time.Duration) time.Duration {
	return 0
}

func (c *Connection) Logger(l modbus.Logger) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.logical = l
}

func (c *Connection) prepare() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.delay()
	c.logger.Logger(c.logical)
}

func (c *Connection) ReadCoils(address, quantity uint16) ([]byte, error) {
	c.prepare()
	return c.ModbusClient().ReadCoils(address, quantity)
}

func (c *Connection) WriteSingleCoil(address, value uint16) ([]byte, error) {
	c.prepare()
	return c.ModbusClient().WriteSingleCoil(address, value)
}

func (c *Connection) ReadInputRegisters(address, quantity uint16) ([]byte, error) {
	c.prepare()
	return c.ModbusClient().ReadInputRegisters(address, quantity)
}

func (c *Connection) ReadHoldingRegisters(address, quantity uint16) ([]byte, error) {
	c.prepare()
	return c.ModbusClient().ReadHoldingRegisters(address, quantity)
}

func (c *Connection) WriteSingleRegister(address, value uint16) ([]byte, error) {
	c.prepare()
	return c.ModbusClient().WriteSingleRegister(address, value)
}

func (c *Connection) WriteMultipleRegisters(address, quantity uint16, value []byte) ([]byte, error) {
	c.prepare()
	return c.ModbusClient().WriteMultipleRegisters(address, quantity, value)
}

func (c *Connection) ReadDiscreteInputs(address, quantity uint16) (results []byte, err error) {
	c.prepare()
	return c.ModbusClient().ReadDiscreteInputs(address, quantity)
}

func (c *Connection) WriteMultipleCoils(address, quantity uint16, value []byte) (results []byte, err error) {
	c.prepare()
	return c.ModbusClient().WriteMultipleCoils(address, quantity, value)
}

func (c *Connection) ReadWriteMultipleRegisters(readAddress, readQuantity, writeAddress, writeQuantity uint16, value []byte) (results []byte, err error) {
	c.prepare()
	return c.ModbusClient().ReadWriteMultipleRegisters(readAddress, readQuantity, writeAddress, writeQuantity, value)
}

func (c *Connection) MaskWriteRegister(address, andMask, orMask uint16) (results []byte, err error) {
	c.prepare()
	return c.ModbusClient().MaskWriteRegister(address, andMask, orMask)
}

func (c *Connection) ReadFIFOQueue(address uint16) (results []byte, err error) {
	c.prepare()
	return c.ModbusClient().ReadFIFOQueue(address)
}
