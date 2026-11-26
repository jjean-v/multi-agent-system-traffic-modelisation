package obj

import (
	"slices"
	"strings"
)

type Tag string

// Matching hierarchique façon GameplayTags (ex: "RoadSign.TrafficLight" match "RoadSign")
func (t Tag) Matches(other Tag) bool {
	if t == other {
		return true
	}
	return strings.HasPrefix(string(t), string(other)+".")
}

type TagContainer struct {
	tags []Tag
}

func NewTagContainer(tags ...Tag) TagContainer {
	tc := TagContainer{}
	for _, t := range tags {
		tc.Add(t)
	}
	return tc
}

func (tc *TagContainer) Add(tag Tag) {
	if !tc.HasTagExact(tag) {
		tc.tags = append(tc.tags, tag)
	}
}

func (tc *TagContainer) Remove(tag Tag) {
	i := slices.Index(tc.tags, tag)
	if i != -1 {
		tc.tags = slices.Delete(tc.tags, i, i+1)
	}
}

func (tc *TagContainer) HasTag(tag Tag) bool {
	for _, t := range tc.tags {
		if t.Matches(tag) {
			return true
		}
	}
	return false
}

func (tc *TagContainer) HasTagExact(tag Tag) bool {
	return slices.Contains(tc.tags, tag)
}

func (tc *TagContainer) HasAny(tags ...Tag) bool {
	for _, t := range tags {
		if tc.HasTag(t) {
			return true
		}
	}
	return false
}

func (tc *TagContainer) HasAll(tags ...Tag) bool {
	for _, t := range tags {
		if !tc.HasTag(t) {
			return false
		}
	}
	return true
}

func (tc *TagContainer) GetTags() []Tag {
	return tc.tags
}
