package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

// data, err := ioutil.ReadAll(os.Stdin)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
// 		os.Exit(1)
// 	}

// 	_, err = os.Stdout.Write(data)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error writing to stdout: %v\n", err)
// 		os.Exit(1)
// 	}

func main() {
	params := os.Args[1:]

	if len(params) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		// os.Stdout.Write(data)
		if err != nil {
			os.Exit(1)
		}
		// fmt.Print(string(data))
		_, err = os.Stdout.Write(data)
		if err != nil {
			os.Exit(1)
		}
		return
	}

	for i := range params {
		content, err := ioutil.ReadFile(params[i])
		if err != nil {
			firstPart := "ERROR: open "
			printStr(firstPart, false)
			fileName := params[i]
			printStr(fileName, false)
			secondPart := ": no such file or directory"
			printStr(secondPart, true)
			os.Exit(1)
		}
		printStr(string(content), false)
	}
}

func printStr(s string, escaping bool) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	if escaping {
		z01.PrintRune('\n')
	}
}
