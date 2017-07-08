package api

type MoveBattleStylePreference struct {
	LowHpPreference      int              `json:"low_hp_preference"`
	HighHpPreference     int              `json:"high_hp_preference"`
	MoveBattleStyleProxy NamedAPIResource `json:"move_battle_style"`
	MoveBattleStyle      MoveBattleStyle
}
