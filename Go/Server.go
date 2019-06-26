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
	var a, b, c, sum int
	fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
	scanner.Scan()
	n := numbers(scanner.Text())
	for j := 0; j < len(n); j++ {
		if sum+n[j] <= b {
			sum += n[j]
			c++
		} else {
			break
		}
	}
	fmt.Println(c)

}
