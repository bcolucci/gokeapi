package api

type Machine struct {
	ID                int              `json:"id"`
	ItemProxy         NamedAPIResource `json:"item"`
	Item              Item
	MoveProxy         NamedAPIResource `json:"move"`
	Move              Move
	VersionGroupProxy NamedAPIResource `json:"version_group"`
	VersionGroup      VersionGroup
}
