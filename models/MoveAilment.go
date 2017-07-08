package models

type MoveAilment struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	MovesProxy []NamedAPIResource `json:"moves"`
	Names      []Name             `json:"names"`
	Moves      []Move
}
