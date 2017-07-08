package models

type Machine struct {
	ID                int              `json:"id"`
	ItemProxy         NamedAPIResource `json:"item"`
	MoveProxy         NamedAPIResource `json:"move"`
	VersionGroupProxy NamedAPIResource `json:"version_group"`
	Item              *Item
	Move              *Move
	VersionGroup      *VersionGroup
}
