package main

import (
"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	if a > b{
		a, b = b, a
	}
	for  i:=a+1; i<b+2; i++ {
		fmt.Println(i)
	}
}
