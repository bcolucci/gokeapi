package models

type MoveMetaData struct {
	AilmentProxy  NamedAPIResource `json:"ailment"`
	CategoryProxy NamedAPIResource `json:"category"`
	MinHits       int              `json:"min_hits"`
	MaxHits       int              `json:"max_hits"`
	MinTurns      int              `json:"min_turns"`
	MaxTurns      int              `json:"max_turns"`
	Drain         int              `json:"drain"`
	Healing       int              `json:"healing"`
	CriticalRate  int              `json:"crit_rate"`
	AilmentChance int              `json:"ailment_chance"`
	FlinchChance  int              `json:"flinch_chance"`
	StatChance    int              `json:"stat_chance"`
	Ailment       *MoveAilment
	Category      *MoveCategory
}
