package brackets

const testVersion = 5

var match = map[rune]rune{
	']': '[',
	'}': '{',
	'"': '"',
	')': '(',
}

func pop(stack []rune) (rune, []rune) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func Bracket(in string) (bool, error) {
	stack := []rune{}

	for _, r := range in {
		if r == '[' || r == '(' || r == '{' || r == '"' {
			stack = append(stack, r)
		} else {
			if r == ']' || r == ')' || r == '}' || r == '"' {
				if len(stack) == 0 {
					return false, nil
				}

				el, newStack := pop(stack)
				stack = newStack

				if el != match[r] {
					return false, nil
				}
			}
		}
	}

	if len(stack) > 0 {
		return false, nil
	}

	return true, nil
}
