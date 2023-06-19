package main

import (
	"os"

	"github.com/01-edu/z01"
)

func bubbleSort(params []string) {
	for i := 0; i < len(params); i++ {
		for j := 0; j < len(params); j++ {
			if params[i] < params[j] {
				params[i], params[j] = params[j], params[i]
			}
		}
	}
}

func main() {
	params := os.Args[1:]
	bubbleSort(params)
	for i := range params {
		for _, r := range params[i] {
			z01.PrintRune(r)
		}
		if i != len(params)-1 {
			z01.PrintRune('\n')
		}

	}
	z01.PrintRune('\n')
}
