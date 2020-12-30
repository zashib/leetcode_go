package leetcode_go

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool
	greatestNum := 0
	for _, candy := range candies {
		if greatestNum < candy {
			greatestNum = candy
		}
	}
	for _, candy := range candies {
		if candy+extraCandies >= greatestNum {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
