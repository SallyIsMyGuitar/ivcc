package core

import (
	"math"
	"time"

	"github.com/evcc-io/evcc/util"
)

// Site is the main configuration container. A site can host multiple loadpoints.
type Savings struct {
	log *util.Logger

	startupTime            time.Time // Boot time
	lastUpdateTime         time.Time // Time of last charged value update
	chargedTotal           float64   // Energy charged since startup (Wh)
	chargedSelfConsumption float64   // Self-produced energy charged since startup (Wh)
}

// NewSite creates a Site with sane defaults
func NewSavings() *Savings {
	savings := &Savings{
		log:            util.NewLogger("savings"),
		startupTime:    time.Now(),
		lastUpdateTime: time.Now(),
	}

	return savings
}

func (s *Savings) Since() int {
	return int(time.Since(s.startupTime).Seconds())
}

func (s *Savings) SelfPercentage() float64 {
	if s.chargedTotal == 0 || s.chargedSelfConsumption == 0 {
		return 0
	}
	return 100 / float64(s.chargedTotal) * float64(s.chargedSelfConsumption)
}

func (s *Savings) ChargedTotal() int {
	return int(s.chargedTotal)
}

func (s *Savings) ChargedSelfConsumption() int {
	return int(s.chargedSelfConsumption)
}

func (s *Savings) shareOfSelfProducedEnergy(gridPower float64, pvPower float64, batteryPower float64) float64 {
	batteryDischarge := math.Max(0, batteryPower)
	batteryCharge := math.Min(0, batteryPower) * -1
	pvConsumption := math.Min(pvPower, pvPower+gridPower-batteryCharge)

	gridImport := math.Max(0, gridPower)
	selfConsumption := math.Max(0, batteryDischarge+pvConsumption+batteryCharge)

	selfPercentage := 100 / (gridImport + selfConsumption) * selfConsumption

	if math.IsNaN(selfPercentage) {
		return 0
	}

	return selfPercentage
}

func (s *Savings) Update(gridPower float64, pvPower float64, batteryPower float64, chargePower float64, mockNow ...time.Time) {
	now := time.Now()
	if len(mockNow) == 1 {
		now = mockNow[0]
	}

	selfPercentage := s.shareOfSelfProducedEnergy(gridPower, pvPower, batteryPower)

	updateDuration := now.Sub(s.lastUpdateTime)

	// assuming the charge power was constant over the duration -> rough estimate
	addedEnergy := updateDuration.Hours() * chargePower

	s.chargedTotal += addedEnergy
	s.chargedSelfConsumption += addedEnergy * (selfPercentage / 100)
	s.lastUpdateTime = now

	s.log.DEBUG.Printf("%.1fkWh charged since %s", s.chargedTotal/1000, time.Since(s.startupTime).Round(time.Second))
	s.log.DEBUG.Printf("%.1fkWh own energy (%.1f%%)", s.chargedSelfConsumption/1000, s.SelfPercentage())
}
