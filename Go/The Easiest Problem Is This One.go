package main

import (
	"fmt"
	"strconv"
)

func sumDigits(num string) int {
	var r int
	for _, v := range num {
		r  = r + int(v - '0')
	}
	return r

}

func main() {
	for {
		var n string
		fmt.Scanln(&n)
		x, _ := strconv.Atoi(n)
		if x == 0 {
			break
		} else {
			inputSum := sumDigits(n)
			for i:=11; i<100000; i++ {
				z := i * x
				if sumDigits(strconv.Itoa(z)) == inputSum {
					fmt.Println(i)
					break
				}
			}
		}

	}
}