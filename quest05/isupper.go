package piscine

func IsUpper(s string) bool {
	cnt := 0

	for _, r := range s {
		if r >= 65 && r <= 90 {
			cnt++
		}
	}
	return cnt == len(s)
}
