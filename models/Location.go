package models

type Location struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Names       []Name                `json:"names"`
	RegionProxy NamedAPIResource      `json:"region"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	AreasProxy  []NamedAPIResource    `json:"areas"`
	Areas       []LocationArea
	Region      *Region
}
