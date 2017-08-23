package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

func Encode(in string) string {
	runes := []rune{}
	for _, r := range in {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			runes = append(runes, unicode.ToLower(r))
		}
	}
	pt := string(runes)

	r, c := setDimensions(len(pt))

	newRows := make([]string, c)

	limit := r

	if r == c {
		limit -= 1
	}

	for i := 0; i <= limit; i++ {
		newRows[i] = newRow(i, c, pt)
	}

	return strings.Join(newRows, " ")
}

func setDimensions(l int) (int, int) {
	n := float64(l)

	sq := math.Sqrt(n)
	r := math.Trunc(sq)

	var c float64
	if r == sq {
		c = r
	} else {
		c = r + 1
	}

	return int(r), int(c)

}

func newRow(i int, c int, pt string) string {
	out := []byte{}
	for idx := i; idx < len(pt); idx += c {
		out = append(out, pt[idx])
	}

	return string(out)
}
