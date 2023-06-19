package piscine

func IsNumeric(s string) bool {
	cnt := 0
	for _, r := range s {
		if r >= 48 && r <= 57 {
			cnt++
		}
	}
	return cnt == len(s)
}
