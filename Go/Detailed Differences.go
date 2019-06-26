package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var s1, s2 string
		fmt.Scanln(&s1)
		fmt.Scanln(&s2)
		fmt.Println(s1)
		fmt.Println(s2)
		for j := 0; j < len(s1); j++ {
			if s1[j] == s2[j] {
				fmt.Print(".")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println("\n")
	}
}
