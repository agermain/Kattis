package main

import (
	"fmt"
)

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func main() {
	var h int
	var s string
	var r int
	fmt.Scanln(&h, &s)
	dom := []int{11, 4, 3, 20, 10, 14, 0, 0}
	nonDom := []int{11, 4, 3, 2, 10, 0, 0, 0}
	suits := []string{"A", "K", "Q", "J", "T", "9", "8", "7"}
	for i:=0; i<h*4; i++ {
		var card string
		fmt.Scanln(&card)
		if string(card[1]) == s {
			r += dom[indexOf(string(card[0]), suits)]
		} else {
			r += nonDom[indexOf(string(card[0]), suits)]
		}
	}
	fmt.Println(r)
}
