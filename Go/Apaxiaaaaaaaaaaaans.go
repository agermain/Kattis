package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scanln(&n)
	var last string
	for _, char := range n {
		if string(char) != last {
			fmt.Printf("%v", char)
			last = string(char)
		}
	}
}
