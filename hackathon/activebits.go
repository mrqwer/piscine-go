package piscine

func ActiveBits(n int) int {
	res := 0
	for n > 0 {
		res += (n % 2)
		n /= 2
	}
	return res
}
