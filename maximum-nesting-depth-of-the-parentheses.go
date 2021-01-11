package leetcode_go

func maxDepth(s string) int {
	parOpen := "("
	parClose := ")"
	depthCurrent := 0
	resultDepth := 0
	for _, char := range s {
		if string(char) == parOpen {
			depthCurrent++
			if depthCurrent > resultDepth {
				resultDepth = depthCurrent
			}
		} else if string(char) == parClose {
			depthCurrent--
		}
	}
	return resultDepth
}
