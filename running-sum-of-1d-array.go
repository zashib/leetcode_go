package leetcode_go

func runningSum(nums []int) []int {
	var result []int
	for i, num := range nums {
		if i == 0 {
			result = append(result, num)
		} else {
			result = append(result, result[i-1]+nums[i])
		}
	}
	return result
}
