package piscine

func CountIf(f func(string) bool, tab []string) int {
	cnt := 0
	for _, v := range tab {
		if f(v) {
			cnt++
		}
	}
	return cnt
}
