package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func contains(sl []string, s string) bool {
	for _, v := range sl {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	w := scanner.Text()
	scanner.Scan()
	e := toInt(scanner.Bytes())
	var sl [][]string
	for i := 0; i < e; i++ {
		scanner.Scan()
		sl = append(sl, strings.Split(scanner.Text(), " "))
	}
	scanner.Scan()
	p := toInt(scanner.Bytes())
	for i := 0; i < p; i++ {
		scanner.Scan()
		s := scanner.Text()
		b := false
		for j := 0; j < len(sl); j++ {
			if contains(sl[j], s) && contains(sl[j], w) {
				b = true
				break
			}
		}
		if b {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
