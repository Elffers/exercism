package pythagorean

// Use this type definition,
//
//    type Triplet [3]int
//
// and implement two functions,
//
//    Range(min, max int) []Triplet
//    Sum(p int) []Triplet
//
// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
//
// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
//
// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.

const testVersion = 1

type Triplet [3]int

func Range(min, max int) []Triplet {
	out := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				t := Triplet{a, b, c}
				if isPythag(t) {
					out = append(out, t)
				}

			}
		}
	}
	return out
}

func Sum(p int) []Triplet {
	out := []Triplet{}
	for _, t := range Range(1, p) {
		a, b, c := t[0], t[1], t[2]
		if a + b + c == p {
				out = append(out, t)
		}
	}

	return out
}

func isPythag(t Triplet) bool {
	a, b, c := t[0], t[1], t[2]
	return a*a+b*b == c*c
}
