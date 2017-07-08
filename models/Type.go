package models

type Type struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	DamageRelations      TypeRelations         `json:"damage_relations"`
	GameIndices          []GenerationGameIndex `json:"game_indicies"`
	GenerationProxy      NamedAPIResource      `json:"generation"`
	MoveDamageClassProxy NamedAPIResource      `json:"move_damage_class"`
	Names                []Name                `json:"names"`
	Pokemon              []TypePokemon         `json:"pokemon"`
	MovesProxy           []NamedAPIResource    `json:"moves"`
	Moves                []Move
	Generation           *Generation
	MoveDamageClass      *MoveDamageClass
}
