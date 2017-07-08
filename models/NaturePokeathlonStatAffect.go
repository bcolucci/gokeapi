package models

type NaturePokeathlonStatAffect struct {
	MaxChange   int              `json:"max_change"`
	NatureProxy NamedAPIResource `json:"nature"`
	Nature      *Nature
}
