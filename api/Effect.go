package api

type Effect struct {
	Effect        string           `json:"effect"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      Language
}
