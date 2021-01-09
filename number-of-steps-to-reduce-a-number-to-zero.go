package leetcode_go

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		count++
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
	}
	return count
}
