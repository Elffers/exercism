package grains

import (
	"fmt"
	"math"
)

const testVersion = 1

func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 1, fmt.Errorf("foo")
	}
	grains := math.Pow(2, float64(s-1))
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
