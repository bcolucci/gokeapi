package api

type TypeRelations struct {
	NoDamageToProxy       []NamedAPIResource `json:"no_damage_to"`
	HalfDamageToProxy     []NamedAPIResource `json:"half_damage_to"`
	DoubleDamageToProxy   []NamedAPIResource `json:"double_damage_to"`
	NoDamageFromProxy     []NamedAPIResource `json:"no_damage_from"`
	HalfDamageFromProxy   []NamedAPIResource `json:"half_damage_from"`
	DoubleDamageFromProxy []NamedAPIResource `json:"double_damage_from"`
	NoDamageTo            []Type
	HalfDamageTo          []Type
	DoubleDamageTo        []Type
	NoDamageFrom          []Type
	HalfDamageFrom        []Type
	DoubleDamageFrom      []Type
}
