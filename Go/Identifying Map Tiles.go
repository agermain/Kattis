package main

import (
	"fmt"
	"math"
)

func main() {
	var a string
	var x, y float64
	fmt.Scanln(&a)
	var s = len(a)
	for i := 0; i < s; i++ {
		if a[i] == '0' {
		}
		if a[i] == '1' {
			x += math.Pow(2, float64(s)-float64(i)-1)
		}
		if a[i] == '2' {
			y += math.Pow(2, float64(s)-float64(i)-1)
		}
		if a[i] == '3' {
			x += math.Pow(2, float64(s)-float64(i)-1)
			y += math.Pow(2, float64(s)-float64(i)-1)
		}
	}
	fmt.Printf("%d %d %d", s, int64(x), int64(y))
}
