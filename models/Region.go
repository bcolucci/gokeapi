package models

type Region struct {
	ID                  int                `json:"id"`
	Name                string             `json:"name"`
	LocationsProxy      []NamedAPIResource `json:"locations"`
	MainGenerationProxy NamedAPIResource   `json:"main_generation"`
	Names               []Name             `json:"names"`
	PokedexesProxy      []NamedAPIResource `json:"pokedexes"`
	VersionGroupProxy   []NamedAPIResource `json:"version_groups"`
	VersionGroup        []VersionGroup
	Locations           []Location
	Pokedexes           []Pokedex
	MainGeneration      *Generation
}
