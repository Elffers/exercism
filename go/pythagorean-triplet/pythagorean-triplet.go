package pythagorean

const testVersion = 1

type Triplet [3]int

func Range(min, max int) []Triplet {
	out := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					out = append(out, Triplet{a, b, c})
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
		if a+b+c == p {
			out = append(out, t)
		}
	}

	return out
}
