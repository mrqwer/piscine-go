package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	for i := len(params) - 1; i >= 0; i-- {
		for _, r := range params[i] {
			z01.PrintRune(r)
		}
		if i != 0 {
			z01.PrintRune('\n')
		}

	}
	z01.PrintRune('\n')
}
