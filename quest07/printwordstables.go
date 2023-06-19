package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i, s := range a {
		for _, r := range s {
			// if i != len(s)-1 {
			// 	z01.PrintRune('\n')
			// }
			z01.PrintRune(r)
		}
		if i != len(a) {
			z01.PrintRune('\n')
		}

	}
}
