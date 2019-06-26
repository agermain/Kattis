package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	fmt.Scanln(&a)
	for {
		m := make(map[int]int)
		var n1, sl1, sl2 []int
		for i := 0; i < a; i++ {
			var d int
			fmt.Scanln(&d)
			n1 = append(n1, d)
			sl1 = append(sl1, d)
		}
		for i := 0; i < a; i++ {
			var d int
			fmt.Scanln(&d)
			sl2 = append(sl2, d)
		}

		sort.Ints(sl1)
		sort.Ints(sl2)

		for i := 0; i < a; i++ {
			m[sl1[i]] = sl2[i]
		}
		for i := 0; i < a; i++ {
			fmt.Printf("%d\n", m[n1[i]])
		}
		fmt.Scanln(&a)
		if a == 0 {
			break
		}
		fmt.Println()
	}
}
