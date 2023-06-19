package main

import (
	"os"

	"github.com/01-edu/z01"
)

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

func IsNumeric(s string) bool {
	cnt := 0
	for _, r := range s {
		if r >= 48 && r <= 57 {
			cnt++
		}
	}
	return cnt == len(s)
}

func main() {
	params := os.Args[1:]
	if len(params) == 0 {
		return
	}
	upper := false
	if params[0] == "--upper" {
		upper = true
	}

	i := 0
	if upper {
		i++
	}
	for ; i < len(params); i++ {
		if isValid(params[i]) {
			var r rune = rune(Atoi(params[i])) + 96
			if upper {
				r = r - 32
			}
			z01.PrintRune(r)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func isValid(s string) bool {
	if IsNumeric(s) && Atoi(s) >= 1 && Atoi(s) <= 26 {
		return true
	}
	return false
}

// A-Z 65-90
// a-z 97-122
