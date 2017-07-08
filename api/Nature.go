package api

type Nature struct {
	ID                         int              `json:"id"`
	Name                       string           `json:"name"`
	DecreasedStatproxy         NamedAPIResource `json:"decreased_stat"`
	DecreasedStat              Stat
	IncreasedStatproxy         NamedAPIResource `json:"increased_stat"`
	IncreasedStat              Stat
	HatesFlavorProxy           NamedAPIResource `json:"hates_flavor"`
	HatesFlavor                BerryFlavor
	LikesFlavorProxy           NamedAPIResource `json:"likes_flavor"`
	LikesFlavor                BerryFlavor
	PokeathlonStatChanges      []NatureStatChange          `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences"`
	Names                      []Name                      `json:"names"`
}
