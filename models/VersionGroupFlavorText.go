package models

type VersionGroupFlavorText struct {
	Text          string           `json:"text"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      *Language
}
