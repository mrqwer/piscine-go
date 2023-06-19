package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i, v := range strs {
		res += string(v)
		if i != len(strs)-1 {
			res += sep
		}
	}
	return res
}
