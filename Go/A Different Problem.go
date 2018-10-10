package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var a, b int64
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			break
		}
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
		fmt.Println(abs(a-b))
	}
}
