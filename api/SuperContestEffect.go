package api

type SuperContestEffect struct {
	ID                int                `json:"id"`
	Appeal            int                `json:"appeal"`
	FlavorTextEntries []FlavorText       `json:"flavor_text_entries"`
	MovesProxy        []NamedAPIResource `json:"moves"`
	Moves             []Move
}
