package api

type Generation struct {
	ID                  int                `json:"id"`
	Name                string             `json:"name"`
	AbilitiesProxy      []NamedAPIResource `json:"abilities"`
	Names               []Name             `json:"names"`
	MainRegionProxy     NamedAPIResource   `json:"main_region"`
	MovesProxy          []NamedAPIResource `json:"moves"`
	PokemonSpeciesProxy []NamedAPIResource `json:"pokemon_species"`
	TypesProxy          []Type             `json:"types"`
	VersionGroupsProxy  []NamedAPIResource `json:"version_groups"`
	Abilities           []Ability
	VersionGroups       []VersionGroup
	Moves               []Move
	PokemonSpecies      []PokemonSpecies
	Types               []Type
	MainRegion          *Region
}
