package models

type NatureStatChange struct {
	MaxChange           int              `json:"max_change"`
	PokeathlonStatProxy NamedAPIResource `json:"pokeathlon_stat"`
	PokeathlonStat      *PokeathlonStat
}
