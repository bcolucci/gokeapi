package api

type Item struct {
	ID                  int              `json:"id"`
	Name                string           `json:"name"`
	Cost                int              `json:"cost"`
	FlingPower          int              `json:"fling_power"`
	FlingEffectProxy    NamedAPIResource `json:"fling_effect"`
	FlingEffect         ItemFlingEffect
	AttributesProxy     []NamedAPIResource `json:"attributes"`
	Attributes          []ItemAttribute
	Category            ItemCategory             `json:"category"`
	EffectEntries       []VerboseEffect          `json:"effect_entries"`
	FlavorTextEntries   []VersionGroupFlavorText `json:"flavor_text_entries"`
	GameIndices         []GenerationGameIndex    `json:"game_indices"`
	Names               []Name                   `json:"names"`
	Sprites             ItemSprites              `json:"sprites"`
	HeldByPokemon       []ItemHolderPokemon      `json:"held_by_pokemon"`
	BabyTriggerForProxy APIResource              `json:"baby_trigger_for"`
	BabyTriggerFor      EvolutionChain
	Machines            []MachineVersionDetail `json:"machines"`
}
