package api

type ItemAttribute struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	ItemsProxy   []NamedAPIResource `json:"items"`
	Items        []Item
	Names        []Name        `json:"names"`
	Descriptions []Description `json:"description"`
}
