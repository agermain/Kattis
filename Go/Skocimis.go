package main

import (
	"bufio"
	"fmt"
	"os"
)

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var a, b, c int64
	fmt.Sscanf(scanner.Text(), "%d %d %d ", &a, &b, &c)
	fmt.Println(Max(b-a, c-b) - 1)
}
