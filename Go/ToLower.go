package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)
func HasUpperCase(str string) bool {
	for i, v := range str {
		if i != 0 && string(v) == strings.ToUpper(string(v)) {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, m, c int
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)
	for i:=0; i<n; i++ {
		solved := true
		for i:=0; i<m; i++ {
			scanner.Scan()
			if HasUpperCase(scanner.Text()) == false {
				solved = false
			}
		}
		if solved {
			c++
		}
	}
	fmt.Println(c)
}
