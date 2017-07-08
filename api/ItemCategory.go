package api

type ItemCategory struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	ItemsProxy  []NamedAPIResource `json:"items"`
	Items       []Item
	Names       []Name           `json:"names"`
	PocketProxy NamedAPIResource `json:"pocket"`
	Pocket      ItemPocket
}
