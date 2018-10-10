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
	n := toInt(scanner.Bytes())
	for i:=0; i<n; i++ {
		scanner.Scan()
		var a, b int
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
		for j:=0; j<b; j++ {
			scanner.Scan()
		}
		fmt.Println(a-1)
	}
}