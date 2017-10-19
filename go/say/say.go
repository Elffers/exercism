package say

import (
	"fmt"
	"math"
	"strings"
)

const testVersion = 1

var magnitude = []struct {
	val  uint64
	word string
}{
	{1e18, "quintillion"},
	{1e15, "quadrillion"},
	{1e12, "trillion"},
	{1e9, "billion"},
	{1e6, "million"},
	{1000, "thousand"},
	{100, "hundred"},
}

var ones = map[uint64]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens = []string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

func Say(num uint64) string {
	if num == 0 {
		return "zero"
	}

	result := ""

	if num < 0 || num > math.MaxUint64 {
		return result
	}

	i := 0

	for i < len(magnitude) {
		n := magnitude[i]
		tally := uint64(0)
		for num >= n.val {
			tally += 1
			num -= n.val
		}

		if tally > 20 {
			result += fmt.Sprintf("%v %v ", hundreds(tally), n.word)
		} else if tally > 0 {
			result += fmt.Sprintf("%v %v ", ones[tally], n.word)
		}
		i += 1
	}
	if num > 0 {
		result += fmt.Sprintf("%v", twoDigit(num))
	}

	return strings.TrimSpace(result)
}

func hundreds(num uint64) string {
	hundreds := num / 100
	rem := num % 100

	result := ""

	if hundreds > 0 {
		result += fmt.Sprintf("%v %v", ones[hundreds], "hundred")
	}

	if rem > 0 {
		result += fmt.Sprintf(" %v", twoDigit(rem))
	}
	return strings.TrimSpace(result)
}

func twoDigit(num uint64) string {
	if num < 20 {
		return ones[num]
	}

	result := ""

	ten := num / 10
	one := num % 10

	result += tens[ten]

	if one > 0 {
		result = fmt.Sprintf("%v-%v", result, ones[one])
	}

	return result
}
