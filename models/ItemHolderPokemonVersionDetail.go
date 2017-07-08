package models

type ItemHolderPokemonVersionDetail struct {
	Rarity       string           `json:"rarity"`
	VersionProxy NamedAPIResource `json:"version"`
	Version      *Version
}
