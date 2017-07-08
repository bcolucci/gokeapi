package api

type MoveTarget struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Descriptions []Description      `json:"description"`
	MovesProxy   []NamedAPIResource `json:"moves"`
	Moves        []Move
	Names        []Name `json:"names"`
}
