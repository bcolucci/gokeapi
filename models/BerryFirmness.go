package models

type BerryFirmness struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	BerriesProxy []NamedAPIResource `json:"berries"`
	Names        []Name             `json:"names"`
	Berries      []Berry
}
