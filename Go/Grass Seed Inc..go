package main

import (
	"fmt"
)

func main() {
	var a, r float64
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	for i := 0; i < b; i++ {
		var w, l float64
		fmt.Scanln(&w, &l)
		r = r + (w * l)
	}
	fmt.Print(r * a)
}
