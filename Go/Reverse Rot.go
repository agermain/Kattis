package main

import (
	"bufio"
	"os"
	"fmt"

)
func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func reverseRot(n int, s string) string {
	sym := "ABCDEFGHIJKLMNOPQRSTUVWXYZ_."
	m := make(map[string]int)
	for i:=0; i<len(sym); i++ {
		m[string(sym[i])] = i
	}
	s = Reverse(s)
	var r string
	for _, v := range s {
		ind := m[string(v)]
		ix := (ind + n) % len(sym)
		r += string(sym[ix])
	}
	return r
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var n int
		var s string
		if scanner.Text() == "0" {
			break
		} else {
			fmt.Sscanf(scanner.Text(), "%d %v", &n, &s)
			fmt.Println(reverseRot(n, s))
		}
	}
}
