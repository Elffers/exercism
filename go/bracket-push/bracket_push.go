package brackets

const testVersion = 5

var match = map[rune]rune{
	']': '[',
	'}': '{',
	')': '(',
}

func Bracket(in string) (bool, error) {
	stack := ""
	for _, r := range in {
		if r == '[' || r == '(' || r == '{' {
			stack += string(r)
		}

		if r == ']' || r == ')' || r == '}' {
			if len(stack) == 0 {
				return false, nil
			}

			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if rune(popped) != match[r] {
				return false, nil
			}
		}
	}

	return len(stack) == 0, nil
}
