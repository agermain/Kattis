package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hidden(a, b string) string {
	for i := 0; i < len(b); i++ {
		if strings.Contains(a, string(b[i])) {
			if b[i] != a[0] {
				return "FAIL"
			}
			a = strings.Replace(a, string(b[i]), "", 1)
		}
	}
	if len(a) == 0 {
		return "PASS"
	} else {
		return "FAIL"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var a, b string
	fmt.Sscanf(scanner.Text(), "%v %v", &a, &b)
	fmt.Println(hidden(a, b))

}
