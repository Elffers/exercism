package anagram

import (
	"strings"
)

const testVersion = 2

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
	if a == b || len(a) != len(b) {
		return false
	}

	for _, r := range a {
		if strings.Count(a, string(r)) != strings.Count(b, string(r)) {
			return false
		}
	}
	return true
}
