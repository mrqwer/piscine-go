package piscine

func ToUpper(s string) string {
	newS := ""
	for i, r := range s {
		if r >= 97 && r <= 122 {
			newS += string(s[i] - 32)
		} else {
			newS += string(s[i])
		}
	}
	return newS
}
