package main

import "fmt"

func main() {
	var k, q, r, b, kn, p int
	fmt.Scanf("%d %d %d %d %d %d", &k, &q, &r, &b, &kn, &p )
	fmt.Printf("%d %d %d %d %d %d", 1-k, 1-q, 2-r, 2-b, 2-kn, 8-p)
}