package lsproduct

import "fmt"

const testVersion = 5

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) || span < 0 {
		return 0, fmt.Errorf("Invalid input")
	}

	max := 0

	for i := 0; i <= len(digits)-span; i++ {
		product := 1

		for _, d := range digits[i : i+span] {
			if d < '0' || d > '9' {
				return 0, fmt.Errorf("Not a valid number")
			}
			product *= int(d - '0')
		}

		if product > max {
			max = product
		}
	}

	return max, nil
}
