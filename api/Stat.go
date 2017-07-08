package api

type Stat struct {
	ID                   int                  `json:"id"`
	Name                 string               `json:"name"`
	GameIndex            int                  `json:"game_index"`
	IsBattleOnly         bool                 `json:"is_battle_only"`
	AffectingMoves       MoveStatAffectSets   `json:"affecting_moves"`
	AffectingNatures     NatureStatAffectSets `json:"affecting_natures"`
	CharacteristicsProxy []APIResource        `json:"characteristics"`
	Characteristics      []Characteristic
	MoveDamageClassProxy NamedAPIResource `json:"move_damage_class"`
	MoveDamageClass      MoveDamageClass
	Names                []Name `json:"names"`
}
