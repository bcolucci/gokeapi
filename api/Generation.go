package api

type Generation struct {
	ID                  int                `json:"id"`
	Name                string             `json:"name"`
	AbilitiesProxy      []NamedAPIResource `json:"abilities"`
	Abilities           []Ability
	Names               []Name           `json:"names"`
	MainRegionProxy     NamedAPIResource `json:"main_region"`
	MainRegion          Region
	MovesProxy          []NamedAPIResource `json:"moves"`
	Moves               []Move
	PokemonSpeciesProxy []NamedAPIResource `json:"pokemon_species"`
	PokemonSpecies      []PokemonSpecies
	TypesProxy          []Type `json:"types"`
	Types               []Type
	VersionGroupsProxy  []NamedAPIResource `json:"version_groups"`
	VersionGroups       []VersionGroup
}
