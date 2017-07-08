package api

type MoveFlavorText struct {
	FlavorText        string           `json:"flavor_text"`
	LanguageProxy     NamedAPIResource `json:"language"`
	Language          Language
	VersionGroupProxy NamedAPIResource `json:"version_group"`
	VersionGroup      VersionGroup
}
