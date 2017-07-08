package models

type MoveTarget struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Descriptions []Description      `json:"description"`
	MovesProxy   []NamedAPIResource `json:"moves"`
	Names        []Name             `json:"names"`
	Moves        []Move
}
