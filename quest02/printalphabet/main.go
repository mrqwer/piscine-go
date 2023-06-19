package main

import "github.com/01-edu/z01"

func main() {
	alph := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range alph {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
