package api

type Description struct {
	Text          string           `json:"description"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      Language
}
