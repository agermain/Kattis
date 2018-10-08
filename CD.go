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

func binarySearch(s []int, value int) int {
	startIndex := 0
	endIndex := len(s) - 1
	for startIndex <= endIndex {
		median := (startIndex + endIndex) / 2
		if s[median] < value {
			startIndex = median + 1
		} else {
			endIndex = median - 1
		}
	}
	if startIndex == len(s) || s[startIndex] != value {
		return -1
	} else {
		return startIndex
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		var n, m int
		var jack []int
		fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)
		if n == 0 && m == 0 {
			break
		}
		c := 0
		for i:=0; i < n; i++ {
			scanner.Scan()
			jack = append(jack, toInt(scanner.Bytes()))
		}
		for i:= 0; i < m; i++ {
			scanner.Scan()
			rs := binarySearch(jack, toInt(scanner.Bytes()))
			if rs != -1 {
				c++
			}
		}
	fmt.Println(c)
	}
}