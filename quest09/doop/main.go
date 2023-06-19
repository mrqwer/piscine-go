package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isValid(params []string) bool {
	if len(params) != 3 {
		return false
	}

	if !isNumeric(params[0]) || !checkOverflow(params[0]) {
		return false
	}

	if !isOperator(params[1]) {
		return false
	}

	if !isNumeric(params[2]) || !checkOverflow(params[2]) {
		return false
	}

	return true
}

func checkOverflow(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Check if the number is negative
	negative := false
	if s[0] == '-' {
		negative = true
		s = s[1:] // Remove the negative sign
	}

	// Check if the number is within the range -2,147,483,648 to 2,147,483,647
	maxValue := "2147483647"
	if negative {
		maxValue = "2147483648"
	}

	if len(s) > len(maxValue) {
		// The number has more digits than the maximum value
		return false
	} else if len(s) < len(maxValue) {
		// The number has fewer digits than the maximum value
		return true
	}

	// Compare each digit with the maximum value
	for i := 0; i < len(s); i++ {
		if s[i] > maxValue[i] {
			// The digit is greater than the maximum value
			return false
		} else if s[i] < maxValue[i] {
			// The digit is smaller than the maximum value
			return true
		}
	}

	// All digits are equal to the maximum value
	return true
}

func isOperator(s string) bool {
	if s == "+" || s == "-" || s == "/" || s == "*" || s == "%" {
		return true
	}
	return false
}

// +, -, /, *, %
func isNumeric(s string) bool {
	cnt := 0
	for _, r := range s {
		if r >= 48 && r <= 57 {
			cnt++
		}
	}
	if s[0] == '-' {
		cnt++
	}
	return cnt == len(s)
}

func BasicAtoi(s string) int {
	sign := 1
	if s[0] == '-' {
		s = s[1:]
		sign = -1
	}
	result := 0
	for _, char := range s {
		num := int(char - '0')
		result = result*10 + num
	}
	return sign * result
}

func basicItoa(n int) string {
	if n == 0 {
		return "0"
	}
	sign := false
	if n < 0 {
		sign = true
		n = -n
	}

	result := ""

	for n > 0 {
		result += string(rune(n%10 + '0'))
		n /= 10
	}

	result = reverseStr(result)
	if sign {
		result = "-" + result
	}
	return result
}

func reverseStr(s string) string {
	res := ""

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

const (
	division = "No division by 0"
	modulo   = "No modulo by 0"
)

func main() {
	params := os.Args[1:]
	if !isValid(params) {
		return
	}

	operand1 := BasicAtoi(params[0])
	operand2 := BasicAtoi(params[2])

	res := ""
	switch params[1] {
	case "+":
		res = basicItoa(operand1 + operand2)
	case "-":
		res = basicItoa(operand1 - operand2)
	case "*":
		res = basicItoa(operand1 * operand2)
	case "/":
		if operand2 == 0 {
			printStr(division)
			return
		}
		res = basicItoa(operand1 / operand2)
	case "%":
		if operand2 == 0 {
			printStr(modulo)
			return
		}
		res = basicItoa(operand1 % operand2)
	}

	printStr(res)
}
