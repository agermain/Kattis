package main

import (
	"bufio"
	"os"
	"fmt"
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
	var r float64
	for i:=0; i<n; i++ {
		var a, b float64
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%f %f", &a, &b)
		r += a*b
	}
	fmt.Println(r)
}