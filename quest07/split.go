package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] && len(s)-i >= len(sep) && s[i:i+len(sep)] == sep {
			if start != i {
				result = append(result, s[start:i])
			}
			i += len(sep) - 1
			start = i + 1
		}
	}
	if start != len(s) {
		result = append(result, s[start:])
	}
	return result
}
