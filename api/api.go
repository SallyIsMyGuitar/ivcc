package api

import "time"

//go:generate mockgen -package mock -destination ../mock/mock.go github.com/andig/evcc/api Charger,Meter,MeterEnergy

// ChargeMode are charge modes modeled after OpenWB
type ChargeMode string

const (
	ModeOff   ChargeMode = "off"
	ModeNow   ChargeMode = "now"
	ModeMinPV ChargeMode = "minpv"
	ModePV    ChargeMode = "pv"
)

// ChargeStatus is the EV's charging status from A to F
type ChargeStatus string

const (
	StatusNone ChargeStatus = ""
	StatusA    ChargeStatus = "A" // Fzg. angeschlossen: nein    Laden möglich: nein
	StatusB    ChargeStatus = "B" // Fzg. angeschlossen:   ja    Laden möglich: nein
	StatusC    ChargeStatus = "C" // Fzg. angeschlossen:   ja    Laden möglich:   ja
	StatusD    ChargeStatus = "D" // Fzg. angeschlossen:   ja    Laden möglich:   ja
	StatusE    ChargeStatus = "E" // Fzg. angeschlossen:   ja    Laden möglich: nein
	StatusF    ChargeStatus = "F" // Fzg. angeschlossen:   ja    Laden möglich: nein
)

// Meter is able to provide current power in W
type Meter interface {
	CurrentPower() (float64, error)
}

// MeterEnergy is able to provide current energy in kWh
type MeterEnergy interface {
	TotalEnergy() (float64, error)
}

// Charger is able to provide current charging status and to enable/disabler charging
type Charger interface {
	Status() (ChargeStatus, error)
	Enabled() (bool, error)
	Enable(enable bool) error
	MaxCurrent(current int64) error
}

// ChargeTimer provides current charge cycle duration
type ChargeTimer interface {
	ChargingTime() (time.Duration, error)
}

// ChargeRater provides charged energy amount in kWh
type ChargeRater interface {
	ChargedEnergy() (float64, error)
}

// Vehicle represents the EV and it's battery
type Vehicle interface {
	Title() string
	Capacity() int64
	ChargeState() (float64, error)
}
