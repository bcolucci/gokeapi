package models

type Nature struct {
	ID                         int                         `json:"id"`
	Name                       string                      `json:"name"`
	DecreasedStatproxy         NamedAPIResource            `json:"decreased_stat"`
	IncreasedStatproxy         NamedAPIResource            `json:"increased_stat"`
	HatesFlavorProxy           NamedAPIResource            `json:"hates_flavor"`
	LikesFlavorProxy           NamedAPIResource            `json:"likes_flavor"`
	PokeathlonStatChanges      []NatureStatChange          `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences"`
	Names                      []Name                      `json:"names"`
	DecreasedStat              *Stat
	IncreasedStat              *Stat
	HatesFlavor                *BerryFlavor
	LikesFlavor                *BerryFlavor
}
