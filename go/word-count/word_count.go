package wordcount

import (
	"strings"
	"unicode"
)

const testVersion = 3

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	out := Frequency{}
	phrase = strings.ToLower(phrase)
	words := wordsFrom(phrase)

	for _, word := range words {
		out[word] += 1
	}
	return out
}

func wordsFrom(phrase string) []string {
	words := []string{}
	split := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	}
	for _, word := range strings.FieldsFunc(phrase, split) {
		word = strings.Trim(word, "'")
		words = append(words, word)
	}
	return words
}
