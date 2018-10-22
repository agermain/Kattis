package main

import (
	"bufio"
	"fmt"
	"os"
)

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var a, b int
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
		if a == 0 && b == 0 {
			break
		} else if a+b == 13 {
			fmt.Println("Never speak again.")
		} else if a < b {
			fmt.Println("Left beehind.")
		} else if a > b {
			fmt.Println("To the convention.")
		} else if a == b {
			fmt.Println("Undecided.")
		}
	}
}
