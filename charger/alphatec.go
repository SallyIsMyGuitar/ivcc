package charger

// LICENSE

// Copyright (c) 2019-2021 andig

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
	"errors"
	"fmt"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/modbus"
)

// https://github.com/evcc-io/evcc/discussions/1965

// Alphatec charger implementation
type Alphatec struct {
	log  *util.Logger
	conn *modbus.Connection
}

const (
	alphatecRegStatus = 0x03
	alphatecRegEnale  = 0x10
	alphatecEnabled   = 1 << 4
)

var AlphatecRegCurrents = []uint16{320, 322, 324}

func init() {
	registry.Add("alphatec", NewAlphatecFromConfig)
}

// NewAlphatecFromConfig creates a Alphatec charger from generic config
func NewAlphatecFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := modbus.Settings{
		ID: 1,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return NewAlphatec(cc.URI, cc.Device, cc.Comset, cc.Baudrate, cc.ID)
}

// NewAlphatec creates Alphatec charger
func NewAlphatec(uri, device, comset string, baudrate int, slaveID uint8) (api.Charger, error) {
	conn, err := modbus.NewConnection(uri, device, comset, baudrate, modbus.TcpFormat, slaveID)
	if err != nil {
		return nil, err
	}

	// if !sponsor.IsAuthorized() {
	// 	return nil, api.ErrSponsorRequired
	// }

	log := util.NewLogger("alpha")
	conn.Logger(log.TRACE)

	wb := &Alphatec{
		log:  log,
		conn: conn,
	}

	// go wb.heartbeat()

	return wb, err
}

// heartbeat implements the api.ChargerEx interface
// func (wb *Alphatec) heartbeat() {
// 	for range time.NewTicker(time.Minute).C {
// 		wb.mu.Lock()
// 		var curr float64
// 		if wb.enabled {
// 			curr = wb.curr
// 		}
// 		wb.mu.Unlock()

// 		if err := wb.setCurrent(curr); err != nil {
// 			wb.log.ERROR.Println("heartbeat:", err)
// 		}
// 	}
// }

// Status implements the api.Charger interface
func (wb *Alphatec) Status() (api.ChargeStatus, error) {
	b, err := wb.conn.ReadHoldingRegisters(alphatecRegStatus, 1)
	if err != nil {
		return api.StatusNone, err
	}

	var res api.ChargeStatus
	switch b[0] {
	case 1:
		res = api.StatusA
	case 2:
		res = api.StatusB
	case 3:
		res = api.StatusC
	default:
		return api.StatusNone, fmt.Errorf("invalid status: %d", b[0])
	}

	return res, nil
}

// Enabled implements the api.Charger interface
func (wb *Alphatec) Enabled() (bool, error) {
	b, err := wb.conn.ReadHoldingRegisters(alphatecEnabled, 1)
	if err != nil {
		return false, err
	}

	return b[0]&alphatecEnabled > 0, nil
}

// Enable implements the api.Charger interface
func (wb *Alphatec) Enable(enable bool) error {
	b := make([]byte, 2)
	if enable {
		b[0] = alphatecEnabled
	}

	_, err := wb.conn.WriteMultipleRegisters(alphatecEnabled, 1, b)

	return err
}

// MaxCurrent implements the api.Charger interface
func (wb *Alphatec) MaxCurrent(current int64) error {
	return errors.New("not implemented")
}
