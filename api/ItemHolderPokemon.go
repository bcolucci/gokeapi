package api

type ItemHolderPokemon struct {
	Pokemon        string                         `json:"pokemon"`
	VersionDetails ItemHolderPokemonVersionDetail `json:"version_details"`
}
