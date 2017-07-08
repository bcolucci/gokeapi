package api

type PokeathlonStat struct {
	ID               int                            `json:"id"`
	Name             string                         `json:"name"`
	Names            []Name                         `json:"names"`
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}
