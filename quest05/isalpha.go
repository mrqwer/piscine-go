package piscine

func alphaCount(s string) int {
	cnt := 0
	for _, r := range s {
		if (r >= 65 && r <= 90) || (r >= 97 && r <= 122) || (r >= 48 && r <= 57) {
			cnt++
		}
	}
	return cnt
}

func IsAlpha(s string) bool {
	return len(s) == alphaCount(s)
}
