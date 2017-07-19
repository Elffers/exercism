package accumulate

const testVersion = 1

func Accumulate(list []string, f func(string) string) []string {
	col := make([]string, len(list))

	for i, str := range list {
		col[i] = f(str)
	}
	return col
}
