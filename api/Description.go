package api

type Description struct {
	Description   string           `json:"description"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      *Language
}
