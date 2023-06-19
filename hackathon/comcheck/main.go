package main

import (
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]
	if len(params) == 0 {
		return
	}

	for _, s := range params {
		if s == "01" || s == "galaxy" || s == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
