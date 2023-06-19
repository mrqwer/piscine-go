package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	cFlag := false
	charCount := -1
	failed := false
	firstNum := 0
	for i, s := range args {
		if s == "-c" {
			cFlag = true
			continue
		}
		if cFlag == true {
			cFlag = false
			charCount = basicAtoi(s)
			firstNum = i + 1
			continue
		}
		content := getContent(s)
		if content == "" {
			failed = true
			continue
		}
		if charCount < 0 {
			charCount = len(content)
		}
		num := len(content) - charCount
		if len(content)-charCount < 0 {
			num = 0
		}
		if firstNum != i {
			fmt.Printf("\n")
		}
		fmt.Printf("==> %s <==\n%s", s, content[num:])
	}
	if failed {
		os.Exit(1)
	}
}

func getContent(filePath string) string {
	f, e := os.Open(filePath)
	if e != nil {
		fmt.Printf("%s\n", e.Error())
		return ""
	}
	stat, _ := f.Stat()
	bytes := make([]byte, stat.Size())
	f.Read(bytes)
	f.Close()
	return string(bytes)
}

func basicAtoi(s string) int {
	result := 0
	for _, char := range s {
		num := int(char - '0')
		result = result*10 + num
	}
	return result
}
