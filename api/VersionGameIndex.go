package api

type VersionGameIndex struct {
	GameIndex    int              `json:"game_index"`
	VersionProxy NamedAPIResource `json:"version"`
	Version      Version
}
