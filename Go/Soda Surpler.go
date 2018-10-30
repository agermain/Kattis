package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var a, b, c int
	fmt.Sscanf(scanner.Text(), "%d %d %d", &a, &b, &c)
	a += b
	var r int
	for a >= c {
		r += a / c
		a = a%c + a/c
	}
	fmt.Println(r)
}
