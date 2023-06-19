package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func byteToRunes(bytes []byte) []rune {
	s := ""

	for _, b := range bytes {
		s += string(rune(b))
	}
	return []rune(s)
}

func main() {
	params := os.Args[1:]
	if len(params) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(params) != 1 {
		fmt.Println("Too many arguments")
		return
	}

	content, _ := ioutil.ReadFile(params[0])
	contentR := byteToRunes(content)
	fmt.Print(string(contentR))
}
