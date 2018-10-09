package main

import "fmt"

func main() {
	var s string
	var c int
	fmt.Scanln(&s)
	n := [3]int{
		0,
		0,
		0,
	}
	for _, d := range s {
		if string(d) == "T" {
			n[0]++
		} else if string(d) == "C" {
			n[1]++
		} else {
			n[2]++
		}
	}
	if n[0] <= n[1] && n[0] <= n[2] {
		c = n[0]
	} else if n[1] <= n[2] {
		c = n[1]
	} else {
		c = n[2]
	}
	fmt.Printf("%d", (7*c)+n[0]*n[0]+n[1]*n[1]+n[2]*n[2])
}
