package queenattack

import (
	"errors"
	"regexp"
)

const testVersion = 2

var validSquareExp = regexp.MustCompile(`^[a-h][1-8]$`)

func validSquare(s string) bool {
	return validSquareExp.MatchString(s)
}

func CanQueenAttack(w, b string) (bool, error) {

	if !(validSquare(b) && validSquare(w)) {
		return false, errors.New("Not a valid square")
	}

	if b == w {
		return false, errors.New("Same square")
	}

	y1 := int(w[0])
	y2 := int(b[0])
	x1 := int(w[1])
	x2 := int(b[1])

	if (x1 - x2) != 0 {
		slope := (y1 - y2) / (x1 - x2)
		if slope == 1 || slope == -1 {
			return true, nil
		}
	}

	if b[0] == w[0] || b[1] == w[1] {
		return true, nil
	}

	return false, nil
}
