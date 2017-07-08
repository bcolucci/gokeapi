package api

type Location struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Names       []Name           `json:"names"`
	RegionProxy NamedAPIResource `json:"region"`
	Region      Region
	GameIndices []GenerationGameIndex `json:"game_indices"`
	AreasProxy  []NamedAPIResource    `json:"areas"`
	Areas       []LocationArea
}
