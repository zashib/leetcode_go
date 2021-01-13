package leetcode_go

func uniqueMorseRepresentations(words []string) int {
	morses := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	conversions := make(map[string]bool)
	for _, s := range words {
		convertion := ""
		runes := []rune(s)
		for _, r := range runes {
			index := r - 'a'
			convertion += morses[index]
		}
		conversions[convertion] = true
	}
	return len(conversions)
}
