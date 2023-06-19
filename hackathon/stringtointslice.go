package piscine

func StringToIntSlice(str string) []int {
	var res []int

	for _, v := range str {
		res = append(res, int(v))
	}
	return res
}
