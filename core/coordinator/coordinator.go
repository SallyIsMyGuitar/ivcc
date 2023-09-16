package coordinator

import (
	"sync"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/loadpoint"
	"github.com/evcc-io/evcc/util"
)

// Coordinator coordinates vehicle access between loadpoints
type Coordinator struct {
	mu       sync.Mutex
	log      *util.Logger
	vehicles []api.Vehicle
	tracked  map[api.Vehicle]loadpoint.API
}

// New creates a coordinator for a set of vehicles
func New(log *util.Logger) *Coordinator {
	return &Coordinator{
		log:     log,
		tracked: make(map[api.Vehicle]loadpoint.API),
	}
}

// SetVehicles sets the list of vehicles, mainly for testing
func (c *Coordinator) SetVehicles(vehicles []api.Vehicle) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.vehicles = vehicles
}

// GetVehicles returns the list of all vehicles
func (c *Coordinator) GetVehicles() []api.Vehicle {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.vehicles
}

func (c *Coordinator) Add(vehicle api.Vehicle) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.vehicles = append(c.vehicles, vehicle)
}

func (c *Coordinator) Delete(vehicle api.Vehicle) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, v := range c.vehicles {
		if v == vehicle {
			c.vehicles = append(c.vehicles[:i], c.vehicles[i+1:]...)

			if o, ok := c.tracked[vehicle]; ok {
				o.SetVehicle(nil)
			}
			delete(c.tracked, vehicle)

			break
		}
	}
}

func (c *Coordinator) acquire(owner loadpoint.API, vehicle api.Vehicle) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if o, ok := c.tracked[vehicle]; ok && o != owner {
		o.SetVehicle(nil)
	}
	c.tracked[vehicle] = owner
}

func (c *Coordinator) release(vehicle api.Vehicle) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.tracked, vehicle)
}

// availableDetectibleVehicles is the list of vehicles that are currently not
// associated to another loadpoint and have a status api that allows for detection
func (c *Coordinator) availableDetectibleVehicles(owner loadpoint.API) []api.Vehicle {
	var res []api.Vehicle

	c.mu.Lock()
	defer c.mu.Unlock()

	for _, vv := range c.vehicles {
		// status api available
		if _, ok := vv.(api.ChargeState); ok {
			// available or associated to current loadpoint
			if o, ok := c.tracked[vv]; o == owner || !ok {
				res = append(res, vv)
			}
		}
	}

	return res
}

// identifyVehicleByStatus finds active vehicle by charge state
func (c *Coordinator) identifyVehicleByStatus(available []api.Vehicle) api.Vehicle {
	var res api.Vehicle

	c.mu.Lock()
	defer c.mu.Unlock()

	for _, vehicle := range available {
		if vs, ok := vehicle.(api.ChargeState); ok {
			status, err := vs.Status()
			if err != nil {
				c.log.ERROR.Println("vehicle status:", err)
				continue
			}

			c.log.DEBUG.Printf("vehicle status: %s (%s)", status, vehicle.Title())

			// vehicle is plugged or charging, so it should be the right one
			if status == api.StatusB || status == api.StatusC {
				if res != nil {
					c.log.WARN.Println("vehicle status: >1 matches, giving up")
					return nil
				}

				res = vehicle
			}
		}
	}

	return res
}
