package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string]string)
	for scanner.Scan() {
		var a, b string
		if scanner.Text() == "" {
			break
		} else {
			fmt.Sscanf(scanner.Text(), "%v %v", &a, &b)
			m[b] = a
		}
	}
	for scanner.Scan() {
		if val, ok := m[scanner.Text()]; ok {
			fmt.Println(val)
		} else {
			fmt.Println("eh")
		}
	}
}
