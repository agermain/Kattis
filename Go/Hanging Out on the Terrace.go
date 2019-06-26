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
	scanner.Scan()
	var n, m, count, current int
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)
	current = 0
	for i := 0; i < m; i++ {
		scanner.Scan()
		var a string
		var b int
		fmt.Sscanf(scanner.Text(), "%v %d", &a, &b)
		if a == "enter" {
			if current+b <= n {
				current += b
			} else {
				count++
			}
		} else {
			current -= b
		}
	}
	fmt.Println(count)
}
