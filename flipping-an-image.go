package leetcode_go

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		start, end := 0, len(row)-1
		for start < end {
			row[start], row[end] = row[end]^1, row[start]^1
			start++
			end--
		}
		if start == end {
			row[start] = row[start] ^ 1
		}
	}
	return A
}
