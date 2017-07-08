package api

type AbilityPokemon struct {
	IsHidden     bool             `json:"is_hidden"`
	Slot         int              `json:"slot"`
	PokemonProxy NamedAPIResource `json:"pokemon"`
	Pokemon      Pokemon
}
