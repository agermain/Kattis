package main

import (
	"fmt"
	"math"
)
func main() {
	var n int
	var w, h float64
	fmt.Scanln(&n, &w, &h)
	for i:=0; i<n; i++ {
		var c float64
		fmt.Scanln(&c)
		max := math.Sqrt(w*w+h*h)
		if c > max {
			fmt.Printf("NE\n")
		} else {
			fmt.Printf("DA\n")
		}
	}
}