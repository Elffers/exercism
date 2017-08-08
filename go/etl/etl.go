package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for score, letters := range input {
		for _, letter := range letters {
			key := strings.ToLower(letter)
			output[key] = score
		}
	}
	return output
}
