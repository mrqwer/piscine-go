package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	pairs := make([]rune, 2)

	for pairs[0] < 99 {
		pairs[1] = pairs[0] + 1
		for pairs[1] < 100 {
			z01.PrintRune(rune(pairs[0]/10) + '0')
			z01.PrintRune(rune(pairs[0]%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(pairs[1]/10) + '0')
			z01.PrintRune(rune(pairs[1]%10) + '0')

			if pairs[0] == 98 && pairs[1] == 99 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			pairs[1]++
		}
		pairs[0]++
	}
}
