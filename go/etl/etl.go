package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		for _, letter := range v {
			newKey := strings.ToLower(letter)
			_, ok := output[newKey]

			if !ok {
				output[newKey] = k
			}
		}
	}
	return output

}
