package leetcode_go

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, jewel := range jewels {
		count += strings.Count(stones, string(jewel))
	}
	return count
}
