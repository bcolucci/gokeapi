package models

type EncounterVersionDetails struct {
	Rate         int              `json:"rate"`
	VersionProxy NamedAPIResource `json:"version"`
	Version      *Version
}
