package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		var k float64
		fmt.Scanln(&k)
		fmt.Println(int(math.Pow(2, k) - 1))
	}
}
