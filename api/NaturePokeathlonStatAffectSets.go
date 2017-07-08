package api

type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect `json:"increase"`
	Decrease []NaturePokeathlonStatAffect `json:"decrease"`
}
