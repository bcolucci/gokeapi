package models

type MoveStatChange struct {
	Change    int              `json:"change"`
	StatProxy NamedAPIResource `json:"stat"`
	Stat      *Stat
}
