package piscine

import "github.com/01-edu/z01"

func printCharByChar(arr []rune, n, last int) {
	i := 0
	if arr[0] == rune(last+'0') {
		for i < n {
			z01.PrintRune(arr[i])
			i++
		}
		z01.PrintRune('\n')
		return
	}
	for i < n {
		z01.PrintRune(arr[i])
		i++
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
}

func PrintCombN(n int) {
	initial := make([]rune, n)
	final := make([]rune, n)
	flag := 0
	base := 0

	for i := 0; i < n; i++ {
		initial[i] = rune('0' + i)
		final[i] = rune('0' + (10 - n) + i)
	}

	printCharByChar(initial, n, int(final[0]-'0'))

	for initial[0] != final[0] {
		flag = n - 1
		for initial[flag] == final[flag] {
			flag--
		}
		base = int(initial[flag]) - '0' + 1

		for flag < n {
			initial[flag] = rune(base + '0')
			flag++
			base++
		}
		printCharByChar(initial, n, int(final[0]-'0'))
	}
}
