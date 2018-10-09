package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	r := a*(b+1)
	for i:=0; i<b; i++{
		fmt.Scanln(&c)
		r -= c
	}
	fmt.Println(r)
}