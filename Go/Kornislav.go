package main

import "fmt"
import "sort"

func main() {
	var a, b, c, d int
	fmt.Scanln(&a, &b, &c, &d)
	ints := []int{a, b, c, d}
	sort.Ints(ints)
	fmt.Printf("%d", ints[0]*ints[2])
}
