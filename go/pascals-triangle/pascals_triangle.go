package pascal

const testVersion = 1

func Triangle(size int) [][]int {
	triangle := make([][]int, size)
	for i := 0; i < size; i++ {
		triangle[i] = Row(i)
	}
	return triangle
}

func Row(n int) []int {
	row := make([]int, n+1)

	for i := 0; i <= n/2; i++ {
		row[i], row[n-i] = comb(n, i), comb(n, i)
	}
	return row
}

// For row n, column k of Pascal's triangle, the value can be
// represented as the combination: n!/(n-k)!k!
func comb(n, k int) int {
	out := 1
	div := 1
	for i := 1; i <= k; i++ {
		out *= n
		n--
		div *= i
	}

	return out / div
}

// func Triangle(n int) [][]int {
// 	tri := make([][]int, n)
// 	for row := 0; row < n; row++ {
// 		tri[row] = make([]int, row+1)
// 		tri[row][0] = 1
// 		for col := 1; col < row; col++ {
// 			tri[row][col] = tri[row-1][col-1] + tri[row-1][col]
// 		}
// 		tri[row][row] = 1
// 	}
// 	return tri
// }
