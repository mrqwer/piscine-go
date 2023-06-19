package piscine

func Index(s string, toFind string) int {
	if len(toFind) > len(s) {
		return -1
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		match := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
