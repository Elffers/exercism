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

	if IsYelling(in) {
		return "Whoa, chill out!"
	}

	if strings.HasSuffix(in, "?") {
		return "Sure."
	}

	return "Whatever."

}

func IsYelling(in string) bool {
	return in == strings.ToUpper(in) && strings.ContainsAny(in, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
