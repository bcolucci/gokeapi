package models

type Name struct {
	Name          string           `json:"name"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      *Language
}
