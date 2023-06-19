package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	// if str == "" {
	// 	return map[string]int{"": 1}
	// }
	m := make(map[string]int)
	splittedWords := splitWhiteSpaces(str)
	// spaces := countSpaces(str)
	// if spaces != 0 {
	// 	m[""] = spaces
	// }

	for _, s := range splittedWords {
		m[s]++
	}

	return m
}

func splitWhiteSpaces(s string) []string {
	var words []string
	index := 0
	for i, char := range s {
		if char == ' ' || char == '\n' || char == '\t' {
			words = append(words, s[index:i])
			index = i + 1
		}
	}
	words = append(words, s[index:])
	return words
}
