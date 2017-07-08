package api

type GenerationGameIndex struct {
	Index           int              `json:"game_index"`
	GenerationProxy NamedAPIResource `json:"generation"`
	Generation      *Generation
}
