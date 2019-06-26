package main

import (
	"fmt"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func main() {
	var r []int
	for i := 0; i < 10; i++ {
		var a int
		fmt.Scanln(&a)
		v := a % 42
		if !contains(r, v) {
			r = append(r, v)
		}
	}
	fmt.Println(len(r))
}
