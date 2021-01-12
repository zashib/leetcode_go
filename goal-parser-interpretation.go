package leetcode_go

import "strings"

func interpret(command string) string {
	var result []string
	chars := strings.Split(command, "")
	for i := 0; i < len(command); i++ {
		if chars[i] == "(" {
			if chars[i+1] == ")" {
				result = append(result, "o")
				i++
				continue
			}
			result = append(result, "al")
			i += 2
			continue
		} else if chars[i] == "G" {
			result = append(result, "G")
		}
	}
	return strings.Join(result, "")
}
