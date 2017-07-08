package api

type ContestComboDetail struct {
	UseBeforeProxy NamedAPIResource `json:"use_before"`
	UseBefore      Move
	UseAfterProxy  NamedAPIResource `json:"use_after"`
	UserAfter      Move
}
