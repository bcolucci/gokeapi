package api

type VersionGroup struct {
	ID                    int                `json:"id"`
	Name                  string             `json:"name"`
	Order                 int                `json:"order"`
	GenerationProxy       NamedAPIResource   `json:"generation"`
	MoveLearnMethodsProxy []NamedAPIResource `json:"move_learn_methods"`
	PokedexesProxy        []NamedAPIResource `json:"pokedexes"`
	RegionsProxy          []NamedAPIResource `json:"regions"`
	VersionsProxy         []NamedAPIResource `json:"versions"`
	MoveLearnMethods      []MoveLearnMethod
	Pokedexes             []Pokedex
	Regions               []Region
	Versions              []Version
	Generation            *Generation
}
