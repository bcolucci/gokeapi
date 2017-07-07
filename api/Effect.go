package api

type Effect struct {
	Text          string           `json:"effect"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      Language
}
