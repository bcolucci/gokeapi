package api

type MoveDamageClass struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Descriptions []Description      `json:"descriptions"`
	MovesProxy   []NamedAPIResource `json:"moves"`
	Names        []Name             `json:"names"`
	Moves        []Move
}
