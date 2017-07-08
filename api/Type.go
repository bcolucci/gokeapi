package api

type Type struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	DamageRelations      TypeRelations         `json:"damage_relations"`
	GameIndices          []GenerationGameIndex `json:"game_indicies"`
	GenerationProxy      NamedAPIResource      `json:"generation"`
	Generation           Generation
	MoveDamageClassProxy NamedAPIResource `json:"move_damage_class"`
	MoveDamageClass      MoveDamageClass
	Names                []Name             `json:"names"`
	Pokemon              []TypePokemon      `json:"pokemon"`
	MovesProxy           []NamedAPIResource `json:"moves"`
	Moves                []Move
}
