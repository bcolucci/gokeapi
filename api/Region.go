package api

type Region struct {
	ID                  int                `json:"id"`
	Name                string             `json:"name"`
	LocationsProxy      []NamedAPIResource `json:"locations"`
	Locations           []Location
	MainGenerationProxy NamedAPIResource `json:"main_generation"`
	MainGeneration      Generation
	Names               []Name             `json:"names"`
	PokedexesProxy      []NamedAPIResource `json:"pokedexes"`
	Pokedexes           []Pokedex
	VersionGroupProxy   []NamedAPIResource `json:"version_groups"`
	VersionGroup        []VersionGroup
}
