package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	for len(a) != len(b) {
		if len(a) < len(b) {
			a = "0" + a
		} else {
			b = "0" + b
		}
	}
	for i:=len(a)-1; i>=0; i-- {
		t1, _ := strconv.Atoi(a[i:i+1])
		t2, _:= strconv.Atoi(b[i:i+1])
		if t1 == t2 {
			continue
		} else if t1 < t2 {
			a = a[0:i] + a[i+1:]
		} else {
			b = b[0:i] + b[i+1:]
		}
	}
	if a == "" {
		fmt.Printf("YODA\n")
	} else {
		r, _ := strconv.Atoi(a)
		fmt.Printf("%d\n", r)
	}
	if b == "" {
		fmt.Printf("YODA\n")
	} else {
		r, _ := strconv.Atoi(b)
		fmt.Printf("%d\n", r)
		}
}