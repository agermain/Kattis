package main

import (
	"bufio"
	"fmt"
	"os"
)
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var a, b int
	fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
	if a == 0 && b == 0{
		fmt.Println("Not a moose")
	} else if a == b {
		fmt.Printf("Even %d", a+b)
	} else {
		fmt.Printf("Odd %d", Max(a, b)*2)
	}
}