package romannumerals

import (
	"fmt"
)

const testVersion = 3

var numerals = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num >= 4000 {
		return "", fmt.Errorf("Invalid range")
	}

	out := ""
	unit := 10 // keeps track of which roman numeral symbol to use

	for num > 0 {
		// converts each digit per decimal place
		out = convert(num, unit) + out
		num = num / 10
		unit *= 10
	}
	return out, nil
}

// Returns roman numeral based on its value modulo 10
func convert(in int, unit int) string {
	out := ""
	fiveUnit := unit / 2 // keeps track of which V, L, or D symbol to use

	num := in % 10
	q := num / 5
	r := num % 5

	// 4 or 9
	if r == 4 {
		out += numerals[unit/10] // prepends appropriate decimal unit before fiveUnit
		if q == 1 {
			out += numerals[unit] // num = 9
		} else {
			out += numerals[fiveUnit] // num = 4
		}
		return out
	}

	// if num >= 5
	if q == 1 {
		out += numerals[fiveUnit]
	}

	// num = 1-3 or 6-8
	for r > 0 {
		out += numerals[unit/10]
		r--
	}

	return out
}
