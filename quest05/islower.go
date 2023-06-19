package piscine

func IsLower(s string) bool {
	cnt := 0
	for _, r := range s {
		if r >= 97 && r <= 122 {
			cnt++
		}
	}
	return cnt == len(s)
}
