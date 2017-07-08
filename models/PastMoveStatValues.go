package models

type PastMoveStatValues struct {
	Accuracy          int              `json:"accuracy"`
	EffectChance      int              `json:"effect_chance"`
	Power             int              `json:"power"`
	PowerPoints       int              `json:"pp"`
	EffectEntries     []VerboseEffect  `json:"effect_entries"`
	TypeProxy         NamedAPIResource `json:"type"`
	VersionGroupProxy NamedAPIResource `json:"version_group"`
	Type              *Type
	VersionGroup      *VersionGroup
}
