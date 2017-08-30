package sieve

const testVersion = 1

func Sieve(n int) []int {
	primes := []int{}

	// Initialize array of all ints from 2 to n
	nums := []int{}
	for i := 2; i <= n; i++ {
		nums = append(nums, i)
	}

	for len(nums) > 0 {
		// pop off first element of nums and add to primes
		prime := nums[0]
		primes = append(primes, prime)
		sieved := []int{}

		for _, p := range nums {
			if p%prime != 0 {
				sieved = append(sieved, p)
			}
		}
		nums = sieved
	}

	return primes
}
