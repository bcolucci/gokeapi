package models

type NatureStatAffectSets struct {
	IncreaseProxy []NamedAPIResource `json:"increase"`
	DecreaseProxy []NamedAPIResource `json:"descrease"`
	Increase      []Nature
	Decrease      []Nature
}
