package api

type EvolutionTrigger struct {
	ID                  int                `json:"id"`
	Name                string             `json:"name"`
	Names               []Name             `json:"names"`
	PokemonSpeciesProxy []NamedAPIResource `json:"pokemon_species"`
	PokemonSpecies      []PokemonSpecies
}
