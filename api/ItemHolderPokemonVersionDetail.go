package api

type ItemHolderPokemonVersionDetail struct {
	Rarity       string           `json:"rarity"`
	VersionProxy NamedAPIResource `json:"version"`
	Version      Version
}
