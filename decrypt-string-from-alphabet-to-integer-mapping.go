package leetcode_go

import "strconv"

func freqAlphabets(s string) string {
	result := []byte{}
	i := len(s) - 1
	for i >= 0 {
		if s[i] == '#' {
			n, _ := strconv.Atoi(s[i-2 : i])
			result = append([]byte{byte(n) - 1 + 'a'}, result...)
			i -= 3
		} else {
			n := s[i] - '0'
			result = append([]byte{n - 1 + 'a'}, result...)
			i--
		}
	}
	return string(result)
}
