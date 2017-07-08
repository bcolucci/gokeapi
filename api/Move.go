package api

type Move struct {
	ID                      int              `json:"id"`
	Name                    string           `json:"name"`
	Accuracy                int              `json:"accuracy"`
	EffectChance            int              `json:"effect_chance"`
	PowerPoints             int              `json:"pp"`
	Priority                int              `json:"priority"`
	Power                   int              `json:"power"`
	ContestCombos           ContestComboSets `json:"contest_combos"`
	ContestTypeProxy        NamedAPIResource `json:"contest_type"`
	ContestType             ContestType
	ContestEffectProxy      APIResource `json:"contest_effect"`
	ContestEffect           ContestEffect
	DomageClassProxy        NamedAPIResource `json:"damage_class"`
	DamageClass             MoveDamageClass
	EffectEntries           []VerboseEffect       `json:"effect_entries"`
	EffectChanges           []AbilityEffectChange `json:"effect_changes"`
	FlavorTextEntries       []MoveFlavorText      `json:"move"`
	GenerationProxy         NamedAPIResource      `json:"generation"`
	Generation              Generation
	Machines                []MachineVersionDetail `json:"machines"`
	Meta                    MoveMetaData           `json:"meta"`
	Names                   []Name                 `json:"names"`
	PastMoveStatValues      []PastMoveStatValues   `json:"past_values"`
	StatChanges             []MoveStatChange       `json:"stat_changes"`
	SuperContestEffectProxy APIResource            `json:"super_contest_effect"`
	SuperContest            SuperContestEffect
	TargetProxy             NamedAPIResource `json:"target"`
	Target                  MoveTarget
	TypeProxy               NamedAPIResource `json:"type"`
	Type                    Type
}
