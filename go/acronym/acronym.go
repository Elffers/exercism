package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Abbreviate(in string) string {
	words := strings.FieldsFunc(in, Delimiter)
	var out string
	for _, word := range words {
		out += string(word[0])
	}

	return strings.ToUpper(out)
}

func Delimiter(c rune) bool {
	return unicode.IsSpace(c) || c == '-'
}
