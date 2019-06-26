package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		var s string
		fmt.Scanln(&s)
		if s != "P=NP" {
			sp := strings.Split(s, "+")
			i1, _ := strconv.Atoi(sp[0])
			i2, _ := strconv.Atoi(sp[1])
			fmt.Println(i1 + i2)
		} else {
			fmt.Printf("skipped\n")
		}
	}
}
