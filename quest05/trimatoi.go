package piscine

func TrimAtoi(s string) int {
	sign := 1
	// 48 57

	res := 0
	for _, r := range s {
		if r == '-' {
			if res == 0 {
				sign = -1
			}
		}
		if r >= 48 && r <= 57 {
			res = res*10 + int(r) - 48
		}

	}
	return res * sign
}
