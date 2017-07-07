package api

type BerryFlavor struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	BerriesProxy     []BerryFlavorMap `json:"berries"`
	ContestTypeProxy NamedAPIResource `json:"contest_type"`
	Names            []Name           `json:"names"`
	Berries          []Berry
}
