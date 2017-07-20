package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hey(in string) string {
	in = strings.TrimSpace(in)
	if in == "" {
		return "Fine. Be that way!"
	}

	if IsYelling(in) == true {
		return "Whoa, chill out!"
	}

	if IsQuestion(in) == true {
		return "Sure."
	}

	return "Whatever."

}

func IsYelling(in string) bool {
	words := strings.Fields(in)

	for _, word := range words {
		if IsAllCaps(word) {
			return true
		}
	}
	return false
}

func IsAllCaps(word string) bool {
	last := string(word[len(word)-1])
	if strings.ContainsAny(last, "!") {
		word = word[0 : len(word)-1]
	}

	if word == "OK" || word == "DMV" {
		return false
	}

	for _, char := range word {
		if !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}

func IsQuestion(in string) bool {
	last := in[len(in)-1]
	return string(last) == "?"
}
