package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	var n int64
	fmt.Scanln(&n)
	fn := strconv.FormatInt(n, 2)
	nw, _ := strconv.ParseInt(Reverse(fn), 2, 64)
	fmt.Println(big.NewInt(nw))
}