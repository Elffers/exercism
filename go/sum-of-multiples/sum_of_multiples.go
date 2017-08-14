package summultiples
const testVersion = 2

func SumMultiples(limit int, divisors... int) int {
	multiples := make([]int, 0)

	for _, div := range divisors {
		// maybe concurrently do all divisors?
		for i := 1; i <= limit/div; i++ {
			mult := div * i
			if mult < limit {
				multiples = appendIfUnique(multiples, mult)
			}
		}
	}

	var sum int
	for _, m := range multiples {
		sum += m
	}

	return sum
}

func appendIfUnique(slice []int, i int) []int {
	for _, el := range slice {
		if el == i {
			return slice
		}
	}
	return append(slice, i)
}
