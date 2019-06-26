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
	for i := 0; i < n; i++ {
		scanner.Scan()
	}
}
