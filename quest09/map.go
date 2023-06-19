package piscine

func Map(f func(int) bool, a []int) []bool {
	bools := make([]bool, len(a))

	for i, v := range a {
		bools[i] = f(v)
	}
	return bools
}
