package api

type Ability struct {
	ID                int              `json:"id"`
	Name              string           `json:"name"`
	IsMainSeries      bool             `json:"is_main_series"`
	GenerationProxy   NamedAPIResource `json:"generation"`
	Generation        Generation
	NamesProxy        []NamedAPIResource `json:"names"`
	Names             []Name
	EffectEntries     []VerboseEffect       `json:"effect_entries"`
	EffectChanges     []AbilityEffectChange `json:"effect_changes"`
	FlavorTextEntries []AbilityFlavorText   `json:"flavor_text_entries"`
	Pokemon           []AbilityPokemon      `json:"pokemon"`
}
