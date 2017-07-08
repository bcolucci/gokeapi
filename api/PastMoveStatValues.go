package api

type PastMoveStatValues struct {
	Accuracy          int              `json:"accuracy"`
	EffectChance      int              `json:"effect_chance"`
	Power             int              `json:"power"`
	PowerPoints       int              `json:"pp"`
	EffectEntries     []VerboseEffect  `json:"effect_entries"`
	TypeProxy         NamedAPIResource `json:"type"`
	Type              Type
	VersionGroupProxy NamedAPIResource `json:"version_group"`
	VersionGroup      VersionGroup
}
