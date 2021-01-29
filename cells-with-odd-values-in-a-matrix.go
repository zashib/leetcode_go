package leetcode_go

func oddCells(n int, m int, indices [][]int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	row := make([]int, n)
	col := make([]int, m)
	l := len(indices)

	for i := 0; i < l; i++ {
		x, y := indices[i][0], indices[i][1]
		row[x]++
		col[y]++
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[i][j] += row[i] + col[j]
			if arr[i][j]%2 != 0 {
				result++
			}
		}
	}
	return result
}
