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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var _ = toInt(scanner.Bytes())
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	var b = true
	for i, v := range s {
		if strconv.Itoa(i+1) == v || v == "mumble" {
			continue
		}
		b = false
		break
	}
	if b == false {
		fmt.Println("something is fishy")
	} else {
		fmt.Println("makes sense")
	}
}
