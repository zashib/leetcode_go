package leetcode_go

import "fmt"

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	tmp_digit := 0
	for _, digit := range fmt.Sprint(n) {
		tmp_digit = int(digit - '0')
		product *= tmp_digit
		sum += tmp_digit
	}
	return product - sum
}
