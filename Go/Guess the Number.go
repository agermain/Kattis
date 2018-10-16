package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	upper := 1000
	lower := 1
	for {
		guess := (upper+lower)/2
		fmt.Println(guess)
		scanner.Scan()
		s := scanner.Text()
		if s == "higher" {
			lower = guess + 1
		} else if s == "lower" {
			upper = guess - 1
		} else if s == "correct" {
			break
		}
	}
	return
}