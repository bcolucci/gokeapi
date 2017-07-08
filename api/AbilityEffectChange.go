package api

type AbilityEffectChange struct {
	EffectEntries      []Effect           `json:"effect_entries"`
	VersionGroupsProxy []NamedAPIResource `json:"version_groups"`
	VersionGroups      []VersionGroup
}
