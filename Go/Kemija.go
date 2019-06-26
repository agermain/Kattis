package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	for i := 0; i < len(s); i++ {
		sr := string(s[i])
		if sr == "a" || sr == "e" || sr == "i" || sr == "o" || sr == "u" {
			fmt.Print(sr)
			i += 2
		} else {
			fmt.Print(sr)
		}
	}
}
