package main

import "github.com/01-edu/z01"

func main() {
	digits := "0123456789"
	for _, char := range digits {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
