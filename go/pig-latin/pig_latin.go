package igpay

import (
	"regexp"
	"strings"
)

const testVersion = 1

var vowels = []string{
	"a",
	"e",
	"i",
	"o",
	"u",
}

var consonantPrefixes = []string{
	"ch",
	"thr",
	"th",
	"qu",
	"squ",
	"sch",
}

func PigLatin(in string) string {
	out := []string{}
	for _, word := range strings.Fields(in) {
		out = append(out, pigLatinize(word))
	}
	return strings.Join(out, " ")
}

func pigLatinize(in string) string {
	suffix := "ay"
	prefix := in
	if !startsWithVowel(in) {
		prefix, suffix = separate(in)
	}
	word := prefix + suffix
	return word
}

func startsWithVowel(in string) bool {
	y := regexp.MustCompile(`^y[^aeiou]`)
	x := regexp.MustCompile(`^x[^aeiou]`)
	for _, v := range vowels {
		if strings.HasPrefix(in, v) || y.MatchString(in) || x.MatchString(in) {
			return true
		}
	}
	return false
}

func separate(in string) (suffix string, prefix string) {
	for _, cons := range consonantPrefixes {
		if strings.HasPrefix(in, cons) {
			parts := strings.SplitAfter(in, cons)
			prefix, suffix = parts[0]+"ay", parts[1]
			return
		}
	}
	prefix = string(in[0]) + "ay"
	suffix = string(in[1:])
	return
}
