package twofer

import "fmt"

func ShareWith(in string) string {
	name := "you"
	if len(in) != 0 {
		name = in
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
