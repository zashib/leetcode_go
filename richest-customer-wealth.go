package leetcode_go

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, wealths := range accounts {
		sumWealth := 0
		for _, wealth := range wealths {
			sumWealth += wealth
		}
		if sumWealth > maxWealth {
			maxWealth = sumWealth
		}
	}
	return maxWealth
}
