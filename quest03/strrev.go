package piscine

func StrRev(s string) string {
	newS := ""
	for i := len(s) - 1; i >= 0; i-- {
		newS += string(s[i])
	}
	return newS
}
