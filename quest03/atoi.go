package piscine

func avoidLeadingZeros3(s string) string {
	i := 0
	for i < len(s) && s[i] == '0' {
		i++
	}
	s = s[i:]
	return s
}

func ContainsNonDigits3(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return true
		}
	}
	return false
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	s = avoidLeadingZeros3(s)
	if len(s) == 0 || ContainsNonDigits3(s) {
		return 0
	}

	result := 0
	for _, digit := range s {
		result = result*10 + int(digit-'0')
	}
	return sign * result
}
