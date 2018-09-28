package main

import (
	"fmt"
)

func main() {
	var a, c int
	fmt.Scanln(&a)
	for i:=0; i<a; i++ {
		var b int
		fmt.Scanf("%d", &b)
		if b < 0 {
			c += 1
		}
	}
	fmt.Println(c)
}