package api

type MoveLearnMethod struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	Descriptions       []Description      `json:"descriptions"`
	Names              []Name             `json:"names"`
	VersionGroupsProxy []NamedAPIResource `json:"version_groups"`
	VersionGroup       []VersionGroup
}
