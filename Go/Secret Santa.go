package main

import "fmt"

func main() {
	var a int
	var t, v float64
	fmt.Scanln(&a)
	if a > 100000 {
		a = 100000
	}
	t = 1
	v = 1
	for i:=2; i<=a; i++ {
		v *= float64(i)
		if i % 2 == 0 {
			t = t - 1 / v
		}else {
			t = t + 1 / v
		}

	}
 fmt.Println(t)
}
