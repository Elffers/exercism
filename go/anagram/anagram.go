package anagram

import (
	"sort"
	"strings"
)

const testVersion = 2

type sortRunes []rune

func Detect(word string, candidates []string) []string {
	word = strings.ToLower(word)

	sublist := []string{}
	for _, c := range candidates {
		if isAnagram(word, strings.ToLower(c)) {
			sublist = append(sublist, c)
		}
	}
	return sublist
}

func isAnagram(a, b string) bool {
	if a == b {
		return false
	}
	return SortString(a) == SortString(b)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
