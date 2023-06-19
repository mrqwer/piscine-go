package piscine

func AppendRange(min, max int) []int {
	var sl []int
	if min >= max {
		return sl
	}
	for i := min; i < max; i++ {
		sl = append(sl, i)
	}
	return sl
}
