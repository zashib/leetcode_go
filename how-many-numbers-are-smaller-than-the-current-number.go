package leetcode_go

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	currentNum := 0
	for i, num := range nums {
		currentNum = num
		for j, otherNum := range nums {
			if j != i {
				if currentNum > otherNum {
					result[i] += 1
				}
			}
		}
	}
	return result
}
