package api

type ChainLink struct {
	IsBaby           bool             `json:"is_baby"`
	SpeciesProxy     NamedAPIResource `json:"species"`
	Species          PokemonSpecies
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	EvolvesTo        []ChainLink       `json:"evolves_to"`
}
