package bob

import (
	"regexp"
	"strings"
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
	hasLetter, _ := regexp.MatchString("[a-zA-Z]", in)
	return in == strings.ToUpper(in) && hasLetter
}
