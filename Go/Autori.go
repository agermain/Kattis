package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	for _, char := range a {
		if unicode.IsUpper(char) {
			b += string(char)
		}
	}
	fmt.Println(b)
}
