package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var l, d, x, m, sum int
	var n = math.MaxInt64
	fmt.Scanln(&l)
	fmt.Scanln(&d)
	fmt.Scanln(&x)
	for i := l; i <= d; i++ {
		s := strconv.Itoa(i)
		for j := 0; j < len(s); j++ {
			sum += int(s[j] - '0')
		}
		if sum == x {
			if i < n {
				n = i
			}
			if i > m {
				m = i
			}
		}
		sum = 0
	}
	fmt.Printf("%d\n%d", n, m)
}
