package api

type ItemFlingEffect struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	EffectEntries []Effect           `json:"effect_entries"`
	ItemsProxy    []NamedAPIResource `json:"items"`
	Items         []Item
}
