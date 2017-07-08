package api

type Encounter struct {
	MinLevel             int                `json:"min_level"`
	MaxLevel             int                `json:"max_level"`
	Chance               int                `json:"chance"`
	MethodProxy          NamedAPIResource   `json:"method"`
	ConditionValuesProxy []NamedAPIResource `json:"condition_values"`
	ConditionValues      []EncounterConditionValue
	Method               *EncounterMethod
}
