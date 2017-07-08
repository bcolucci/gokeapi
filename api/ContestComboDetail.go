package api

type ContestComboDetail struct {
	UseBeforeProxy NamedAPIResource `json:"use_before"`
	UseAfterProxy  NamedAPIResource `json:"use_after"`
	UseBefore      *Move
	UserAfter      *Move
}
