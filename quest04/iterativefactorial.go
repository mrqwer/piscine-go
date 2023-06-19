package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 10000 {
		return 0
	}

	fact := 1

	for i := 2; i <= nb; i++ {
		fact *= i
	}

	return fact
}
