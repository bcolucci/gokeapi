package models

type PokemonEncounter struct {
	EncounterMethodProxy NamedAPIResource          `json:"encounter_method"`
	VersionDetails       []EncounterVersionDetails `json:"version_details"`
	EncounterMethod      *EncounterMethod
}
