package api

type TypeRelations struct {
	NoDamageToProxy       []NamedAPIResource `json:"no_damage_to"`
	NoDamageTo            []Type
	HalfDamageToProxy     []NamedAPIResource `json:"half_damage_to"`
	HalfDamageTo          []Type
	DoubleDamageToProxy   []NamedAPIResource `json:"double_damage_to"`
	DoubleDamageTo        []Type
	NoDamageFromProxy     []NamedAPIResource `json:"no_damage_from"`
	NoDamageFrom          []Type
	HalfDamageFromProxy   []NamedAPIResource `json:"half_damage_from"`
	HalfDamageFrom        []Type
	DoubleDamageFromProxy []NamedAPIResource `json:"double_damage_from"`
	DoubleDamageFrom      []Type
}
