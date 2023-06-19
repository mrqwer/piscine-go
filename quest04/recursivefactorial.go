package piscine

const maxFactorialInput = 10000

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}

	if nb < 0 {
		return 0
	}
	if nb > maxFactorialInput {
		return 0
	}

	return nb * RecursiveFactorial(nb-1)
}
