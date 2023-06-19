package piscine

func Sqrt(nb int) int {
	res := 0
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			res = i
			break
		}
		if i*i > nb {
			return 0
		}
	}
	return res
}

// sqrt(4) = 2
