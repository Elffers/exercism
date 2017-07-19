package pangram

import "unicode"
import "strings"

const testVersion = 1

func IsPangram(s string) bool {
	seen := make(map[string]bool)
	for _, char := range strings.ToLower(s) {
		if unicode.IsLetter(char) {
			seen[string(char)] = true
		}
	}

	for _, char := range "abcdefghijklmnopqrstuvwxyz" {
		_, ok := seen[string(char)]
		if ok == false {
			return false
		}
	}
	return true
}
