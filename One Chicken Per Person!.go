package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	if m > n && m-n == 1 {
		fmt.Println("Dr. Chaz will have 1 piece of chicken left over!")
	} else if m > n {
		fmt.Printf("Dr. Chaz will have %d pieces of chicken left over!", m-n)
	} else if m < n && m-n == -1 {
		fmt.Println("Dr. Chaz needs 1 more piece of chicken!")
	} else {
		fmt.Printf("Dr. Chaz needs %d more pieces of chicken!", n-m)
	}
}