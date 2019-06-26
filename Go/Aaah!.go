package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l1 := len(scanner.Bytes())
	scanner.Scan()
	l2 := len(scanner.Bytes())
	if l1 >= l2 {
		fmt.Println("go")
	} else {
		fmt.Println("no")
	}
}
