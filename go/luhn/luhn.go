package luhn

import (
	"unicode"
)

const testVersion = 2

func Valid(in string) bool {
	num, ok := sanitize(in)
	if !ok {
		return false
	}

	if len(num) <= 1 {
		return false
	}

	sum := 0
	double := false
	for i := len(num) - 1; i >= 0; i-- {
		digit := int(num[i] - '0')
		if double {
			doubled := digit * 2
			if doubled > 9 {
				doubled -= 9
			}
			sum += doubled
		} else {
			sum += digit
		}
		double = !double
	}

	return sum%10 == 0
}

func sanitize(in string) (string, bool) {
	out := ""
	for _, r := range in {
		if !(unicode.IsNumber(r) || unicode.IsSpace(r)) {
			return "", false
		}
		if unicode.IsNumber(r) {
			out += string(r)
		}
	}
	return out, true
}
