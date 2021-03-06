package series

const testVersion = 2

func All(n int, s string) []string {
	substrings := make([]string, 0)

	for i := 0; i <= len(s)-n; i++ {
		substring := s[i : i+n]
		substrings = append(substrings, substring)

	}
	return substrings
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
