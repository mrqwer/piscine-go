package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	runes := []rune(programName)
	for i := 2; i < len(runes); i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
