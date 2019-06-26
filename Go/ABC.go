package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	var s string
	fmt.Scanln(&a, &b, &c)
	fmt.Scanln(&s)
	var l = []int{a, b, c}
	sort.Ints(l)
	for i := 0; i < 3; i++ {
		if string(s[i]) == "A" {
			fmt.Printf("%d ", l[0])
		} else if string(s[i]) == "B" {
			fmt.Printf("%d ", l[1])
		} else {
			fmt.Printf("%d ", l[2])
		}
	}
}
