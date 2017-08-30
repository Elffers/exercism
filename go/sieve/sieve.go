package sieve

const testVersion = 1

func Sieve(n int) []int {
	primes := []int{}
	sieve := make([]bool, n)

	for p := 2; p < n; p++ {
		if !sieve[p] {
			// element p is prime if value at index p in sieve is false
			primes = append(primes, p)
			for multiple := p * 2; multiple < n; multiple += p {
				// filter out multiples of primes
				sieve[multiple] = true
			}
		}
	}

	return primes
}
