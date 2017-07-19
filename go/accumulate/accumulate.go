package accumulate

const testVersion = 1

func Accumulate(list []string, f func(string) string) []string {
	col := make([]string, len(list))
	for _, str := range list {
		col = append(col, f(str))[1:]
	}
	return col
}
