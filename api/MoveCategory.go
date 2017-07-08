package api

type MoveCategory struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Descriptions []Description      `json:"descriptions"`
	MovesProxy   []NamedAPIResource `json:"moves"`
	Moves        []Move
}
