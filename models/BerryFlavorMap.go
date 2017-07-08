package models

type BerryFlavorMap struct {
	Potency    int              `json:"potency"`
	BerryProxy NamedAPIResource `json:"berry"`
	Berry      *Berry
}
