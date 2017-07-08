package models

type EncounterCondition struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Names       []Name             `json:"names"`
	ValuesProxy []NamedAPIResource `json:"values"`
	Values      []EncounterConditionValue
}
