package models

type ChainLink struct {
	IsBaby           bool              `json:"is_baby"`
	SpeciesProxy     NamedAPIResource  `json:"species"`
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	EvolvesTo        []ChainLink       `json:"evolves_to"`
	Species          *PokemonSpecies
}
