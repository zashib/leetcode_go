package leetcode_go

import "strings"

func countConsistentStrings(allowed string, words []string) int {
	result := 0
	for _, word := range words {
		for i, char := range word {
			if strings.Contains(allowed, string(char)) != true {
				break
			} else if strings.Contains(allowed, string(char)) == true && i == len(word)-1 {
				result++
			}
		}
	}
	return result
}
