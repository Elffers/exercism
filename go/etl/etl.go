package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		for _, letter := range v {
			new_key := strings.ToLower(letter)
			_, ok := output[new_key]

			if !ok {
				output[new_key] = k
			}
		}
	}
	return output

}
