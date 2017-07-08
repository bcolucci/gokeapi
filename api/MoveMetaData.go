package api

type MoveMetaData struct {
	AilmentProxy  NamedAPIResource `json:"ailment"`
	Ailment       MoveAilment
	CategoryProxy NamedAPIResource `json:"category"`
	Category      Move
	MinHits       int `json:"min_hits"`
	MaxHits       int `json:"max_hits"`
	MinTurns      int `json:"min_turns"`
	MaxTurns      int `json:"max_turns"`
	Drain         int `json:"drain"`
	Healing       int `json:"healing"`
	CriticalRate  int `json:"crit_rate"`
	AilmentChance int `json:"ailment_chance"`
	FlinchChance  int `json:"flinch_chance"`
	StatChance    int `json:"stat_chance"`
}
