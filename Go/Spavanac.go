package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	b -= 45
	if 0 > b {
		b += 60
		a--
		if 0 > a {
			a += 24
		}
	}
	fmt.Printf("%d %d", a, b)
}
