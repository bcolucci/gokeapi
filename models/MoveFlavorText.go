package models

type MoveFlavorText struct {
	FlavorText        string           `json:"flavor_text"`
	LanguageProxy     NamedAPIResource `json:"language"`
	VersionGroupProxy NamedAPIResource `json:"version_group"`
	Language          *Language
	VersionGroup      *VersionGroup
}
