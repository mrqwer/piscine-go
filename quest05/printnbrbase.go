package piscine

import "github.com/01-edu/z01"

func reversePrint(rs []rune, sign bool) {
	if sign {
		z01.PrintRune('-')
	}
	for i := len(rs) - 1; i >= 0; i-- {
		z01.PrintRune(rs[i] + '0')
	}
}

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 || containsSignCharacters(base) || !isUniqueCharacters(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < -2147483648 {
		s1 := ""
		temp1 := -nbr
		for temp1 < 0 {
			s1 += string(rune((temp1 % 10) * -1))
			temp1 /= 10
		}

		rs1 := []rune(s1)
		reversePrint(rs1, true)
		return
	}
	if nbr > 2147483647 {
		s2 := ""
		temp2 := nbr
		for temp2 > 0 {
			s2 += string(rune(temp2 % 10))
			temp2 /= 10
		}
		rs2 := []rune(s2)
		reversePrint(rs2, false)
		return
	}

	negative := false
	if nbr < 0 {
		negative = true
		nbr = -nbr
	}
	result := ""
	b := len(base)

	for nbr > 0 {
		remainder := nbr % b
		result = string(base[remainder]) + result
		nbr = nbr / b
	}

	if negative {
		result = "-" + result
	}

	for _, r := range result {
		z01.PrintRune(r)
	}
}

func containsSignCharacters(base string) bool {
	for _, r := range base {
		if r == '+' || r == '-' {
			return true
		}
	}
	return false
}

func isUniqueCharacters(base string) bool {
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}
