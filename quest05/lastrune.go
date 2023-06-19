package piscine

func LastRune(s string) rune {
	rs := []rune(s)
	return rs[len(rs)-1]
}
