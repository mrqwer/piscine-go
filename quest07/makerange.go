package piscine

func MakeRange(min, max int) []int {
	var invalid []int
	if min >= max {
		return invalid
	}
	sl := make([]int, max-min)
	index := 0
	for value := min; value < max; value++ {
		sl[index] = value
		index++
	}
	return sl
}
