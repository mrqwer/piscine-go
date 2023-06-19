package piscine

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}
