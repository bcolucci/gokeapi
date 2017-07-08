package models

type Stat struct {
	ID                   int                  `json:"id"`
	Name                 string               `json:"name"`
	GameIndex            int                  `json:"game_index"`
	IsBattleOnly         bool                 `json:"is_battle_only"`
	AffectingMoves       MoveStatAffectSets   `json:"affecting_moves"`
	AffectingNatures     NatureStatAffectSets `json:"affecting_natures"`
	CharacteristicsProxy []APIResource        `json:"characteristics"`
	MoveDamageClassProxy NamedAPIResource     `json:"move_damage_class"`
	Names                []Name               `json:"names"`
	Characteristics      []Characteristic
	MoveDamageClass      *MoveDamageClass
}
