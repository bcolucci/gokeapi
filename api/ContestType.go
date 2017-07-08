package api

type ContestType struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	BerryFlavorProxy NamedAPIResource `json:"berry_flavor"`
	BerryFlavor      *BerryFlavor
}
