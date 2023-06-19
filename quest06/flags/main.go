package main

import (
	"fmt"
	"os"
)

func orderString(runes []rune) string {
	n := len(runes)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if runes[i-1] > runes[i] {
				runes[i-1], runes[i] = runes[i], runes[i-1]
				swapped = true
			}
		}
		n--
	}

	return string(runes)
}

func extractInsertionValue(s string, short bool) string {
	if short {
		return s[3:]
	}
	return s[9:]
}

func isInsertOption(s string) bool {
	if len(s) <= 9 {
		if len(s) > 3 && s[:2] == "-i" {
			return true
		}
		return false
	}
	return s[:9] == "--insert=" || s[:3] == "-i="
}

func isOrderOption(s string) bool {
	return s == "--order" || s == "-o"
}

func main() {
	params := os.Args[1:]
	if len(params) == 0 || (len(params) == 1 && (params[0] == "--help" || params[0] == "-h")) {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("\t This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
		return
	}

	result := ""
	insertionValues := ""
	order := false
	for _, s := range params {
		if isInsertOption(s) {
			short := false
			if s[0] == '-' && s[1] != '-' {
				short = true
			}
			insertionValues += extractInsertionValue(s, short)
		} else if isOrderOption(s) {
			order = true
		} else {
			result += s
		}
	}
	result += insertionValues
	if order {
		runes := []rune(result)
		result = orderString(runes)
	}
	fmt.Println(result)
}
