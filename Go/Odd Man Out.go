package main

import (
	"bufio"
	"bytes"
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
		scanner.Scan()
		a := bytes.Split(scanner.Bytes(), []byte(" "))
		ints := make([]int, 0)
		for _, b := range a {
			ints = append(ints, toInt(b))
		}
		var bx int
		for j := 0; j < len(ints); j++ {
			bx ^= ints[j]
		}
		fmt.Printf("Case #%d: %d\n", i+1, bx)
	}
}
