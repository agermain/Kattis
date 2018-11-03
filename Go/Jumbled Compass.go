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
	var n, m int
	n = toInt(scanner.Bytes())
	scanner.Scan()
	m = toInt(scanner.Bytes())
	if n == m  {
		fmt.Println("0")
	} else if n < m {
		v1 := m-n
		v2 := (360-m)+n
		if v1 <= v2 {
			fmt.Println(v1)
		} else {
			fmt.Println(-v2)
		}
	} else {
		v1 := n-m
		v2 := (360-n)+m
		if v1 < v2 {
			fmt.Println(-v1)
		} else {
			fmt.Println(v2)
		}
	}
}