package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanln(&n)
	for i:=0; i<n; i++ {
		var v, a, x, h1, h2 float64
		fmt.Scanln(&v, &a, &x, &h1, &h2)
		a = a * math.Pi/180.0
		time := x/(v*math.Cos(a))
		f := v * time * math.Sin(a) - 9.81 * time * time / 2
		if f >= h1+1 && f <= h2-1 {
			fmt.Println("Safe")
		} else {
			fmt.Println("Not Safe")
		}
	}
}
