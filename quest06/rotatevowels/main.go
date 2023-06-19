package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(c rune) bool {
	if c == 'a' || c == 'A' || c == 'e' ||
		c == 'E' || c == 'i' || c == 'I' ||
		c == 'o' || c == 'O' || c == 'u' || c == 'U' {
		return true
	}
	return false
}

func reverseVowels(s string) string {
	vowels := []rune{}
	runes := []rune(s)

	for i := range runes {
		if isVowel(runes[i]) {
			vowels = append(vowels, runes[i])
		}
	}

	j := len(vowels) - 1
	for i := range runes {
		if isVowel(runes[i]) {
			runes[i] = vowels[j]
			j--
		}
	}

	return string(runes)
}

func main() {
	params := os.Args[1:]
	if len(params) < 1 {
		z01.PrintRune('\n')
		return
	}

	result := ""
	for i, s := range params {
		result += s
		if i != len(params)-1 {
			result += " "
		}
	}
	result = reverseVowels(result)

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
