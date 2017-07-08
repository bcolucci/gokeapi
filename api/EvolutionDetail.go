package api

type EvolutionDetail struct {
	ItemProxy             NamedAPIResource `json:"item"`
	Item                  Item
	TriggerProxy          NamedAPIResource `json:"trigger"`
	Trigger               EvolutionTrigger
	Gender                int              `json:"gender"`
	HeldItemProxy         NamedAPIResource `json:"held_item"`
	HeldItem              Item
	KnownMoveProxy        NamedAPIResource `json:"known_move"`
	KnownMove             Move
	KnownMoveTypeProxy    NamedAPIResource `json:"known_move_type"`
	KnownMoveType         Type
	LocationProxy         NamedAPIResource `json:"location"`
	Location              Location
	MinLevel              int              `json:"min_level"`
	MinHappiness          int              `json:"min_hapiness"`
	MinBeauty             int              `json:"min_beauty"`
	MinAffection          int              `json:"min_affection"`
	NeedsOverworldRain    bool             `json:"needs_overworld_rain"`
	PartySpeciesProxy     NamedAPIResource `json:"party_species"`
	PartySpecies          PokemonSpecies
	PartyTypeProxy        NamedAPIResource `json:"party_type"`
	PartyType             Type
	RelativePhysicalStats int              `json:"relative_physical_stats"`
	TimeOfDay             string           `json:"time_of_day"`
	TradeSpeciesProxy     NamedAPIResource `json:"trade_species"`
	TradeSpecies          PokemonSpecies
	turn_upside_down      bool `json:"turn_upside_down"`
}
