package piscine

func avoidLeadingZeros2(s string) string {
	i := 0
	for i < len(s) && s[i] == '0' {
		i++
	}
	s = s[i:]
	return s
}

func ContainsNonDigits(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return true
		}
	}
	return false
}

func power2(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func BasicAtoi2(s string) int {
	s = avoidLeadingZeros(s)
	if len(s) == 0 {
		s = "0"
	}
	if ContainsNonDigits(s) {
		return 0
	}
	exp := len(s) - 1

	result := 0
	for _, digit := range s {
		result = result*10 + int(digit-'0')
		exp--
	}
	return result
}
