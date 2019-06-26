package main

import (
	"fmt"
	"math"
)

func compute(h, v float64) int {
	v = v * math.Pi / 180
	return int(math.Ceil(h / math.Sin(v)))
}
func main() {
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Print(compute(a, b))

}
