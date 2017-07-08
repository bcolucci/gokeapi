package api

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	LocationProxy        NamedAPIResource      `json:"location"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
	Location             *Region
}
