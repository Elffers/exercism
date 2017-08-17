package isogram

import "unicode"

const testVersion = 1

func IsIsogram(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		key := unicode.ToLower(r)
		if _, ok := seen[key]; ok {
			return false
		}
		if unicode.IsLetter(r) {
			seen[key] = true
		}
	}
	return true
}
