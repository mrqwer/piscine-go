package piscine

func JumpOver(str string) string {
	var runes []rune
	for index, char := range str {
		if (index+1)%3 == 0 {
			runes = append(runes, char)
		}
	}
	runes = append(runes, '\n')
	return string(runes)
}
