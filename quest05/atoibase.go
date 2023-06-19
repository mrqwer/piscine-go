package piscine

func AtoiBase(s string, base string) int {
	if len(base) < 2 || containsSignCharacters(base) || !isUniqueCharacters(base) {
		return 0
	}

	result := 0
	for i := 0; i < len(s); i++ {
		idx := -1
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				idx = j
				break
			}
		}
		if idx == -1 {
			return 0
		}
		result = result*len(base) + idx
	}
	return result
}
