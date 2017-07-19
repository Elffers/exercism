package hamming

import "fmt"

const testVersion = 6

type LengthError struct {
	Message string
}

func (e LengthError) Error() string {
	return fmt.Sprintf(e.Message)
}

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, LengthError{"lengths are not equal"}
	}

	count := 0

	for i, c := range a {
		if c != rune(b[i]) {
			count++
		}
	}

	return count, nil
}
