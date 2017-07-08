package models

type Move struct {
	ID                      int                    `json:"id"`
	Name                    string                 `json:"name"`
	Accuracy                int                    `json:"accuracy"`
	EffectChance            int                    `json:"effect_chance"`
	PowerPoints             int                    `json:"pp"`
	Priority                int                    `json:"priority"`
	Power                   int                    `json:"power"`
	ContestCombos           ContestComboSets       `json:"contest_combos"`
	ContestTypeProxy        NamedAPIResource       `json:"contest_type"`
	ContestEffectProxy      APIResource            `json:"contest_effect"`
	DomageClassProxy        NamedAPIResource       `json:"damage_class"`
	EffectEntries           []VerboseEffect        `json:"effect_entries"`
	EffectChanges           []AbilityEffectChange  `json:"effect_changes"`
	FlavorTextEntries       []MoveFlavorText       `json:"move"`
	GenerationProxy         NamedAPIResource       `json:"generation"`
	Machines                []MachineVersionDetail `json:"machines"`
	Meta                    MoveMetaData           `json:"meta"`
	Names                   []Name                 `json:"names"`
	PastMoveStatValues      []PastMoveStatValues   `json:"past_values"`
	StatChanges             []MoveStatChange       `json:"stat_changes"`
	SuperContestEffectProxy APIResource            `json:"super_contest_effect"`
	TargetProxy             NamedAPIResource       `json:"target"`
	TypeProxy               NamedAPIResource       `json:"type"`
	ContestType             *ContestType
	ContestEffect           *ContestEffect
	DamageClass             *MoveDamageClass
	Generation              *Generation
	SuperContest            *SuperContestEffect
	Target                  *MoveTarget
	Type                    *Type
}
