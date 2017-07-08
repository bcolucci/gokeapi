package api

type TypePokemon struct {
	Slot         int              `json:"slot"`
	PokemonProxy NamedAPIResource `json:"pokemon"`
	Pokemon      Pokemon
}
