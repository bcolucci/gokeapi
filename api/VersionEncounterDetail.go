package api

type VersionEncounterDetail struct {
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
	VersionProxy     NamedAPIResource `json:"version"`
	Version          *Version
}
