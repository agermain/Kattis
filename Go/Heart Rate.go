package main

import (
	"fmt"
)

func main() {
	var n int
	var b, p float64
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&b, &p)
		fmt.Printf("%.4f %.4f %.4f\n", ((60/p)*b)-(60/p), (60/p)*b, ((60/p)*b)+(60/p))
	}
}
