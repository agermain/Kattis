package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)
	for i:= 0; i<a; i++{
		var a, b, c int64
		fmt.Scanln(&a, &b, &c)
		if (b - c) == a{
			fmt.Printf("does not matter\n")
		}else if b - c > a {
			fmt.Printf("advertise\n")
		}else{
			fmt.Printf("do not advertise\n")
		}
	}

}