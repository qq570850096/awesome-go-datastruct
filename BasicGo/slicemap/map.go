package slicemap

import "strings"

// CountWords 统计字符串中每个单词出现的次数。
func CountWords(text string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(text) {
		m[w]++
	}
	return m
}

// Set 是用 map 模拟的字符串集合。
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

