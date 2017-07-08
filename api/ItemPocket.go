package api

type ItemPocket struct {
	ID              int                `json:"id"`
	Name            string             `json:"name"`
	CategoriesProxy []NamedAPIResource `json:"categories"`
	Names           []Name             `json:"names"`
	Categories      []ItemCategory
}
