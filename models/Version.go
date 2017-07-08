package models

type Version struct {
	ID                int              `json:"id"`
	Name              string           `json:"name"`
	Names             []Name           `json:"names"`
	VersionGroupProxy NamedAPIResource `json:"version_group"`
	VersionGroup      *VersionGroup
}
