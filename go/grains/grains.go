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
	for i := 0; i < 65; i++ {
		s, err := Square(i)
		if err == nil {
			sum += s
		}
	}
	return
}
