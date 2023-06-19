package piscine

func AlphaCount(s string) int {
	cnt := 0
	for _, r := range s {
		if (r >= 65 && r <= 90) || (r >= 97 && r <= 122) {
			cnt++
		}
	}
	return cnt
}
