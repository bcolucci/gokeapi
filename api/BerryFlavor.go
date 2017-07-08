package api

type BerryFlavor struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	Berries          []BerryFlavorMap `json:"berries"`
	ContestTypeProxy NamedAPIResource `json:"contest_type"`
	Names            []Name           `json:"names"`
	ContestType      *ContestType
}
