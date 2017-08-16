package grains

import (
	"fmt"
)

const testVersion = 1

func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 1, fmt.Errorf("foo")
	}
	grains := 1 << uint(s-1)
	return uint64(grains), nil
}

func Total() (sum uint64) {
	s, err := Square(64)

	if err != nil {
		return 0
	}

	return (2 * s) - 1
}
