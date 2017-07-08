package api

type FlavorText struct {
	Text          string           `json:"flavor_text"`
	LanguageProxy NamedAPIResource `json:"language"`
	Language      *Language
}
