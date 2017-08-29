package luhn

import (
	"strings"
)

const testVersion = 2

func Valid(in string) bool {
	in = strings.TrimSpace(in)

	if len(in) <= 1 {
		return false
	}

	sum, double := 0, false

	for i := len(in) - 1; i >= 0; i-- {
		r := in[i]

		if r == ' ' {
			continue
		}

		if r < '0' || r > '9' {
			return false
		}

		digit := int(r - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	return sum%10 == 0
}
