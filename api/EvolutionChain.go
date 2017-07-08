package api

type EvolutionChain struct {
	ID                   int              `json:"id"`
	BabyTriggerItemProxy NamedAPIResource `json:"baby_trigger_item"`
	Chain                ChainLink        `json:"chain"`
	BabyTriggerItem      *Item
}
