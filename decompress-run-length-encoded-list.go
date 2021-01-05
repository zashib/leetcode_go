package leetcode_go

func decompressRLElist(nums []int) []int {
	var result []int
	for i, num := range nums {
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				result = append(result, nums[i+1])
			}
		}
	}
	return result
}
