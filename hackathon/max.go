package piscine

func Max(a []int) int {
	m := a[0]
	for _, v := range a {
		if m < v {
			m = v
		}
	}
	return m
}
