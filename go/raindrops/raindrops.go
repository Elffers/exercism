package raindrops

import "strconv"

const testVersion = 3

func Convert(num int) string {
	var output string
	if num%3 == 0 {
		output += "Pling"
	}
	if num%5 == 0 {
		output += "Plang"
	}
	if num%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		return strconv.Itoa(num)
	} else {
		return output
	}
}
