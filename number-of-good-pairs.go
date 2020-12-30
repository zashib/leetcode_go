package leetcode_go

func numIdenticalPairs(nums []int) int {
	result := 0
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i < j && num1 == num2 {
				result++
			}
		}
	}
	return result
}
