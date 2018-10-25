package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var sum int
	sl := make([]int, n+2)
	for i:=0; i<n; i++ {
		scanner.Scan()
		sl[i] = toInt(scanner.Bytes())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
	for i:=0; i<n; i++ {
			sum += sl[i] + sl[i+1]
			i+=2
	}
	fmt.Println(sum)
}