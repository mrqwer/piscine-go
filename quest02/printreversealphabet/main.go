package main

import "github.com/01-edu/z01"

func main() {
	alph := "abcdefghijklmnopqrstuvwxyz"
	for i := len(alph) - 1; i >= 0; i-- {
		z01.PrintRune(rune(alph[i]))
	}
	z01.PrintRune('\n')
}
