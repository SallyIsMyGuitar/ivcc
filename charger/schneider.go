package charger

// LICENSE

// Copyright (c) 2019-2023 andig

// This module is NOT covered by the MIT license. All rights reserved.

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/modbus"
	"github.com/evcc-io/evcc/util/sponsor"
	"github.com/volkszaehler/mbmd/encoding"
)

// Schneider charger implementation
type Schneider struct {
	conn *modbus.Connection
	curr uint16
}

const (
	schneiderRegEvState       = 1
	schneiderRegCurrents      = 2999
	schneiderRegVoltages      = 3027
	schneiderRegPower         = 3059
	schneiderRegEnergy        = 3203
	schneiderRegSetCommand    = 4001
	schneiderRegSetPoint      = 4004
	schneiderRegChargingTime  = 4007
	schneiderRegSessionEnergy = 4012
)

func init() {
	registry.Add("schneider", NewSchneiderFromConfig)
}

// https://download.schneider-electric.com/files?p_enDocType=Other+technical+guide&p_File_Name=GEX1969300-03.pdf&p_Doc_Ref=GEX1969300

// NewSchneiderFromConfig creates a Schneider charger from generic config
func NewSchneiderFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		modbus.Settings `mapstructure:",squash"`
		Timeout         time.Duration
	}{
		Settings: modbus.Settings{
			ID: 255,
		},
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return NewSchneider(cc.URI, cc.Device, cc.Comset, cc.Baudrate, cc.ID, cc.Timeout)
}

// go:generate go run ../cmd/tools/decorate.go -f decorateSchneider -b *Schneider -r api.Charger -t "api.Meter,CurrentPower,func() (float64, error)" -t "api.PhaseCurrents,Currents,func() (float64, float64, float64, error)"

// NewSchneider creates Schneider charger
func NewSchneider(uri, device, comset string, baudrate int, slaveID uint8, timeout time.Duration) (api.Charger, error) {
	conn, err := modbus.NewConnection(uri, device, comset, baudrate, modbus.Tcp, slaveID)
	if err != nil {
		return nil, err
	}

	if timeout > 0 {
		conn.Timeout(timeout)
	}

	if !sponsor.IsAuthorized() {
		return nil, api.ErrSponsorRequired
	}

	log := util.NewLogger("schneider")
	conn.Logger(log.TRACE)

	wb := &Schneider{
		conn: conn,
		curr: 6,
	}

	return wb, nil
}

// Status implements the api.Charger interface
func (wb *Schneider) Status() (api.ChargeStatus, error) {
	b, err := wb.conn.ReadHoldingRegisters(schneiderRegEvState, 1)
	if err != nil {
		return api.StatusNone, err
	}

	s := binary.BigEndian.Uint16(b)

	switch s {
	case 2, 6:
		return api.StatusA, nil
	case 3, 4, 5, 7:
		return api.StatusB, nil
	case 8, 9:
		return api.StatusC, nil
	default:
		return api.StatusNone, fmt.Errorf("invalid status: %d", s)
	}
}

// Enabled implements the api.Charger interface
func (wb *Schneider) Enabled() (bool, error) {
	b, err := wb.conn.ReadHoldingRegisters(schneiderRegSetPoint, 1)
	if err != nil {
		return false, err
	}

	return binary.BigEndian.Uint16(b) > 0, nil
}

// Enable implements the api.Charger interface
func (wb *Schneider) Enable(enable bool) error {
	var u uint16
	if enable {
		u = wb.curr
	}

	_, err := wb.conn.WriteSingleRegister(schneiderRegSetPoint, u)
	return err
}

// MaxCurrent implements the api.Charger interface
func (wb *Schneider) MaxCurrent(current int64) error {
	_, err := wb.conn.WriteSingleRegister(schneiderRegSetPoint, uint16(current))
	if err == nil {
		wb.curr = uint16(current)
	}

	return err
}

// CurrentPower implements the api.Meter interface
func (wb *Schneider) CurrentPower() (float64, error) {
	b, err := wb.conn.ReadHoldingRegisters(schneiderRegPower, 2)
	if err != nil {
		return 0, err
	}

	return float64(encoding.Float32LswFirst(b)), nil
}

var _ api.MeterEnergy = (*Schneider)(nil)

// TotalEnergy implements the api.MeterEnergy interface
func (wb *Schneider) TotalEnergy() (float64, error) {
	b, err := wb.conn.ReadHoldingRegisters(schneiderRegEnergy, 2)
	if err != nil {
		return 0, err
	}

	return float64(encoding.Float32LswFirst(b)), nil
}

var _ api.ChargeRater = (*Schneider)(nil)

// ChargedEnergy implements the api.MeterEnergy interface
func (wb *Schneider) ChargedEnergy() (float64, error) {
	b, err := wb.conn.ReadHoldingRegisters(schneiderRegSessionEnergy, 2)
	if err != nil {
		return 0, err
	}

	return float64(encoding.Float32LswFirst(b)), nil
}

var _ api.PhaseCurrents = (*Schneider)(nil)

// Currents implements the api.PhaseCurrents interface
func (wb *Schneider) Currents() (float64, float64, float64, error) {
	b, err := wb.conn.ReadHoldingRegisters(schneiderRegCurrents, 6)
	if err != nil {
		return 0, 0, 0, err
	}

	var res []float64
	for i := 0; i < 3; i++ {
		res = append(res, float64(encoding.Float32LswFirst(b[4*i:])))
	}

	return res[0], res[1], res[2], nil
}

var _ api.PhaseVoltages = (*Schneider)(nil)

// Voltages implements the api.PhaseVoltages interface
func (wb *Schneider) Voltages() (float64, float64, float64, error) {
	b, err := wb.conn.ReadHoldingRegisters(schneiderRegVoltages, 6)
	if err != nil {
		return 0, 0, 0, err
	}

	var res []float64
	for i := 0; i < 3; i++ {
		res = append(res, float64(encoding.Float32LswFirst(b[4*i:])))
	}

	return res[0], res[1], res[2], nil
}
