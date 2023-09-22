package stringset

import (
	"fmt"
	"strings"
)

type nothing struct{}

type Set map[string]nothing

func New() Set {
	return make(Set)
}

func NewFromSlice(list []string) Set {
	set := make(Set)
	for _, l := range list {
		set[l] = nothing{}
	}
	return set
}

func (s Set) String() string {
	var b strings.Builder

	_, _ = b.WriteString("{")

	for element := range s {
		_, _ = fmt.Fprintf(&b, "\"%s\", ", element)
	}

	return strings.TrimSuffix(b.String(), ", ") + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = nothing{}
}

func Subset(s1, s2 Set) bool {
	for element := range s1 {
		if _, ok := s2[element]; !ok {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	set := make(Set)

	for element := range s1 {
		if _, ok := s2[element]; ok {
			set.Add(element)
		}
	}

	for element := range s2 {
		if _, ok := s1[element]; ok {
			set.Add(element)
		}
	}

	return set
}

func Difference(s1, s2 Set) Set {
	set := make(Set)

	for element := range s1 {
		if _, ok := s2[element]; !ok {
			set.Add(element)
		}
	}

	return set
}

func Union(s1, s2 Set) Set {
	set := make(Set)

	for s := range s1 {
		set.Add(s)
	}

	for s := range s2 {
		set.Add(s)
	}

	return set
}
