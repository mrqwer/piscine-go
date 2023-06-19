package piscine

func NRune(s string, n int) rune {
	rs := []rune(s)

	if n > len(rs) || n <= 0 {
		return 0
	}

	return rs[n-1]
}
