package api

type Name struct {
	Name          string           `json:"name"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      *Language
}
