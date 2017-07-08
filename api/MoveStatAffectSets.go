package api

type MoveStatAffectSets struct {
	Increase []MoveStatAffect `json:"increase"`
	Decrease []MoveStatAffect `json:"decrease"`
}
