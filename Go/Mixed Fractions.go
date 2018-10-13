package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var a, b int
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
		if a==0 && b==0 {
			break
		} else {
			fmt.Printf("%d %d / %d\n", a/b, a%b, b)
		}
	}

}