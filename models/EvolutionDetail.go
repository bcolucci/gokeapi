package models

type EvolutionDetail struct {
	ItemProxy             NamedAPIResource `json:"item"`
	TriggerProxy          NamedAPIResource `json:"trigger"`
	Gender                int              `json:"gender"`
	HeldItemProxy         NamedAPIResource `json:"held_item"`
	KnownMoveProxy        NamedAPIResource `json:"known_move"`
	KnownMoveTypeProxy    NamedAPIResource `json:"known_move_type"`
	LocationProxy         NamedAPIResource `json:"location"`
	MinLevel              int              `json:"min_level"`
	MinHappiness          int              `json:"min_hapiness"`
	MinBeauty             int              `json:"min_beauty"`
	MinAffection          int              `json:"min_affection"`
	NeedsOverworldRain    bool             `json:"needs_overworld_rain"`
	PartySpeciesProxy     NamedAPIResource `json:"party_species"`
	PartyTypeProxy        NamedAPIResource `json:"party_type"`
	RelativePhysicalStats int              `json:"relative_physical_stats"`
	TimeOfDay             string           `json:"time_of_day"`
	TradeSpeciesProxy     NamedAPIResource `json:"trade_species"`
	TurnUpsideDown        bool             `json:"turn_upside_down"`
	Item                  *Item
	Trigger               *EvolutionTrigger
	HeldItem              *Item
	KnownMove             *Move
	KnownMoveType         *Type
	Location              *Location
	PartySpecies          *PokemonSpecies
	PartyType             *Type
	TradeSpecies          *PokemonSpecies
}
