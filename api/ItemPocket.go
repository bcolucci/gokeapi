package api

type ItemPocket struct {
	ID              int                `json:"id"`
	Name            string             `json:"name"`
	CategoriesProxy []NamedAPIResource `json:"categories"`
	Categories      []ItemCategory
	Names           []Name `json:"names"`
}
