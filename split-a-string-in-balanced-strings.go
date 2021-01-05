package leetcode_go

func balancedStringSplit(s string) int {
	countL := 0
	countR := 0
	resultCount := 0
	for _, char := range s {
		if string(char) == "L" {
			countL++
		} else {
			countR++
		}
		if countL == countR {
			resultCount++
		}
	}
	return resultCount
}
