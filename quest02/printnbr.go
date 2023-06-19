package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	sign := 1
	if n < 0 {
		sign = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		lessTh10 := (n / 10) * sign
		if lessTh10 != 0 {
			PrintNbr(lessTh10)
		}
		z01.PrintRune(rune((n % 10 * sign)) + '0')
	} else {
		z01.PrintRune('0')
	}
}
