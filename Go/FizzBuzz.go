package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)

	for i := 1; i <= c; i++ {
		if i%a == 0 && i%b == 0 {
			fmt.Printf("FizzBuzz\n")
			continue
		} else if i%a == 0 {
			fmt.Printf("Fizz\n")
			continue
		} else if i%b == 0 {
			fmt.Printf("Buzz\n")
			continue
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
