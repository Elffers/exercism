package palindrome

import (
	"fmt"
	"strconv"
)

const testVersion = 1

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func Products(fmin, fmax int) (pmin, pmax Product, e error) {
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			product := i * j

			if palindrome(product) {
				// Check max
				if product > pmax.Product {
					pmax.Product = product
					pmax.Factorizations = [][2]int{{i, j}}
				} else if product == pmax.Product {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
				}

				// Check min
				if pmin.Product == 0 {
					pmin.Product = product
				}
				if product < pmin.Product {
					pmin.Product = product
					pmin.Factorizations = [][2]int{{i, j}}
				} else if product == pmin.Product {
					pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
				}
			}
		}
	}
	if pmin.Product == 0 && pmax.Product == 0 {
		return Product{}, Product{}, fmt.Errorf("no palindromes")
	}
	return pmin, pmax, nil
}

func palindrome(p int) bool {
	// If negative numbers are considered palindromes
	if p < 0 {
		p = -p
	}
	s := strconv.Itoa(p)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
