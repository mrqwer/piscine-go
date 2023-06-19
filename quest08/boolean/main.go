package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func main() {
	params := os.Args[1:]
	lengthOfArg := len(params)
	evenMsg := "I have an even number of arguments"
	oddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) {
		printStr(evenMsg)
	} else {
		printStr(oddMsg)
	}
}
