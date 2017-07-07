package api

type EncounterConditionValue struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	ConditionProxy NamedAPIResource `json:"condition"`
	Condition      EncounterCondition
}
