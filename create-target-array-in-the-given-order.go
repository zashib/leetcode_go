package leetcode_go

func createTargetArray(nums []int, index []int) []int {
	var targetArray []int
	for i, v := range index {
		if v != i {
			targetArray = append(targetArray[:v+1], targetArray[v:]...)
			targetArray[v] = nums[i]

		} else {
			targetArray = append(targetArray, nums[i])
		}
	}
	return targetArray
}
