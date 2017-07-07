package api

type BerryFirmness struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	BerriesProxy []NamedAPIResource `json:"berries"`
	Berries      []Berry
	Names        []Name
}
