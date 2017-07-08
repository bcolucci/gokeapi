package api

type AbilityFlavorText struct {
	FlavorText         string             `json:"flavor_text"`
	VersionGroupsProxy []NamedAPIResource `json:"version_groups"`
	VersionGroups      []VersionGroup
	LanguageProxy      NamedAPIResource `json:"language"`
	Language           Language
}
