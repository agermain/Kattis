package main

import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}
func main() {
	var a, b string
	fmt.Scanln(&a, &b)
	c, _ := strconv.ParseInt(Reverse(a), 10, 0)
	d, _ := strconv.ParseInt(Reverse(b), 10, 0)
	fmt.Printf("%d", Max(c, d))
}
