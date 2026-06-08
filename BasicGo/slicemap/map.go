package slicemap

import "strings"

func CountWords(text string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(text) {
		m[w]++
	}
	return m
}

type Set map[string]struct{}

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(v string) {
	s[v] = struct{}{}
}

func (s Set) Has(v string) bool {
	_, ok := s[v]
	return ok
}

func (s Set) Remove(v string) {
	delete(s, v)
}
