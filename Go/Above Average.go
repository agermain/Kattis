package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := toInt(scanner.Bytes())
	for i := 0; i < n; i++ {
		var n []int
		scanner.Scan()
		n = numbers(scanner.Text())
		var s, c int
		for j := 1; j < len(n); j++ {
			s += n[j]
		}
		var avg float64
		avg = float64(s) / float64(n[0])
		for j := 1; j < len(n); j++ {
			if n[j] > int(avg) {
				c++
			}
		}
		fmt.Printf("%.3f", (float64(c)/float64(len(n)-1))*100)
		fmt.Print("%\n")
	}
}
