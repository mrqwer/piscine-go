package piscine

func ReverseMenuIndex(menu []string) []string {
	res := make([]string, len(menu))

	j := 0
	for i := len(menu) - 1; i >= 0; i-- {
		res[j] = menu[i]
		j++
	}
	// res = menu
	return res
}
