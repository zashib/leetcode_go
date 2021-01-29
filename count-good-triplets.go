package leetcode_go

func countGoodTriplets(arr []int, a int, b int, c int) int {
	if len(arr) < 3 {
		return 0
	}
	result, n := 0, len(arr)
	for i := range arr {
		for j := i + 1; j < n; j++ {
			if intAbs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if intAbs(arr[j]-arr[k]) > b {
					continue
				}
				if intAbs(arr[i]-arr[k]) > c {
					continue
				}
				result++
			}
		}
	}
	return result
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
