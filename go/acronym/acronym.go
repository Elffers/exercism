package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Abbreviate(in string) (out string) {
	words := strings.FieldsFunc(in, Delimiter)
	for _, word := range words {
		out += strings.ToUpper(string(word[0]))
	}

	return
}

func Delimiter(c rune) bool {
	return unicode.IsSpace(c) || c == '-'
}
