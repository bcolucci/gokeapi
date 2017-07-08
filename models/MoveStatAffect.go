package models

type MoveStatAffect struct {
	Change    int              `json:"change"`
	MoveProxy NamedAPIResource `json:"move"`
	Move      *Move
}
