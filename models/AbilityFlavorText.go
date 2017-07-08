package models

type AbilityFlavorText struct {
	FlavorText         string             `json:"flavor_text"`
	VersionGroupsProxy []NamedAPIResource `json:"version_groups"`
	LanguageProxy      NamedAPIResource   `json:"language"`
	VersionGroups      []VersionGroup
	Language           *Language
}
