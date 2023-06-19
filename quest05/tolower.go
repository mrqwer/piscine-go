package piscine

func ToLower(s string) string {
	newS := ""
	for i, r := range s {
		if r >= 65 && r <= 90 {
			newS += string(s[i] + 32)
		} else {
			newS += string(s[i])
		}
	}
	return newS
}
