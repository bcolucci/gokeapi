package api

type Berry struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
	GrowthTime           int              `json:"growth_time"`
	MaxHarvest           int              `json:"max_harvest"`
	NaturalGiftPower     int              `json:"natural_gift_power"`
	NaturalGiftTypeProxy NamedAPIResource `json:"natural_gift_type"`
	Size                 int              `json:"size"`
	Smoothness           int              `json:"smoothness"`
	SoilDryness          int              `json:"soil_dryness"`
	FirmnessProxy        NamedAPIResource `json:"firmness"`
	Flavors              []BerryFlavorMap `json:"flavors"`
	ItemProxy            NamedAPIResource `json:"item"`
	NaturalGiftType      Type
	Firmness             BerryFirmness
	Item                 Item
}
