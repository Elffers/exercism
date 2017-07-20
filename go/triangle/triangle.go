package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if IsTriangle(a, b, c) == false {
		return NaT
	}

	if IsEqu(a, b, c) == true {
		return Equ
	}

	if IsIso(a, b, c) == true {
		return Iso
	}

	if IsSca(a, b, c) == true {
		return Sca
	}
	return NaT
}

func IsTriangle(a, b, c float64) bool {
	if !IsValid(a) || !IsValid(b) || !IsValid(c) {
		return false
	}
	if (a+b) < c || (a+c) < b || (b+c) < a {
		return false
	}
	return true
}
func IsValid(a float64) bool {
	if math.IsNaN(a) {
		return false
	}
	if math.IsInf(a, 0) {
		return false
	}
	if a <= 0 {
		return false
	}
	return true
}

func IsEqu(a, b, c float64) bool {
	return (a == b) && (b == c)
}

func IsIso(a, b, c float64) bool {
	return a == b || a == c || b == c
}

func IsSca(a, b, c float64) bool {
	return !IsEqu(a, b, c) && !IsIso(a, b, c)
}

type Kind uint

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)
