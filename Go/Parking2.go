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
func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := toInt(scanner.Bytes())
	for i := 0; i < n; i++ {
		scanner.Scan()
		ints := make([]int, toInt(scanner.Bytes()))
		scanner.Scan()
		a := bytes.Split(scanner.Bytes(), []byte(" "))
		for i, b := range a {
			ints[i] = toInt(b)
		}
		x, y := MinMax(ints)
		fmt.Println((y - x) * 2)
	}
}
