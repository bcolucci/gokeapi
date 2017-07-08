package models

type ItemCategory struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	ItemsProxy  []NamedAPIResource `json:"items"`
	Names       []Name             `json:"names"`
	PocketProxy NamedAPIResource   `json:"pocket"`
	Items       []Item
	Pocket      *ItemPocket
}
