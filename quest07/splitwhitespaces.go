package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	var word string
	for _, char := range s {
		switch char {
		case ' ', '\t', '\n':
			if word != "" {
				words = append(words, word)
				word = ""
			}
		default:
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
