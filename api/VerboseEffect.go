package api

type VerboseEffect struct {
	Effect        string           `json:"effect"`
	ShortEffect   string           `json:"short_effect"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      *Language
}
