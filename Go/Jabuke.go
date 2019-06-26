package main

import (
	"fmt"
	"math"
)

func area(x1, y1, x2, y2, x3, y3 float64) float64 {
	return math.Abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)) / 2.0)
}

func isInside(x1, y1, x2, y2, x3, y3, x, y float64) bool {
	var A = area(x1, y1, x2, y2, x3, y3)
	var A1 = area(x, y, x2, y2, x3, y3)
	var A2 = area(x1, y1, x, y, x3, y3)
	var A3 = area(x1, y1, x2, y2, x, y)
	return A == A1+A2+A3
}

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	var it, c int
	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)
	fmt.Scanln(&x3, &y3)
	fmt.Scanln(&it)
	for i := 0; i < it; i++ {
		var x, y float64
		fmt.Scanln(&x, &y)
		if isInside(x1, y1, x2, y2, x3, y3, x, y) {
			c += 1
		}
	}
	fmt.Printf("%.1f\n", area(x1, y1, x2, y2, x3, y3))
	fmt.Println(c)
}
