package api

//go:generate enumer -type TariffType -trimprefix TariffType -transform=lower

type TariffType int

const (
	_ TariffType = iota
	TariffTypePrice
	TariffTypeCo2
)
