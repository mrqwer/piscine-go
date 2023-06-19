package piscine

func avoidLeadingZeros(s string) string {
	i := 0
	for i < len(s) && s[i] == '0' {
		i++
	}
	s = s[i:]
	return s
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func BasicAtoi(s string) int {
	s = avoidLeadingZeros(s)
	if len(s) == 0 {
		s = "0"
	}
	exp := len(s) - 1

	result := 0
	for _, digit := range s {
		result = result*10 + int(digit-'0')
		exp--
	}
	return result
}
