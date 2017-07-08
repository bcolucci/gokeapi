package api

type MoveAilment struct {
	ID         int                `json:"id"`
	Name       string             `json:"name"`
	MovesProxy []NamedAPIResource `json:"moves"`
	Moves      []Move
	Names      []Name `json:"names"`
}
