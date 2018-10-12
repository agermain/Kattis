package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var m = map[byte][]byte{
		'a': []byte("@"),
		'b': []byte("8"),
		'c': []byte("("),
		'd': []byte("|)"),
		'e': []byte("3"),
		'f': []byte("#"),
		'g': []byte("6"),
		'h': []byte("[-]"),
		'i': []byte("|"),
		'j': []byte("_|"),
		'k': []byte("|<"),
		'l': []byte("1"),
		'm': []byte("[]\\/[]"),
		'n': []byte("[]\\[]"),
		'o': []byte("0"),
		'p': []byte("|D"),
		'q': []byte("(,)"),
		'r': []byte("|Z"),
		's': []byte("$"),
		't': []byte("']['"),
		'u': []byte("|_|"),
		'v': []byte("\\/"),
		'w': []byte("\\/\\/"),
		'x': []byte("}{"),
		'y': []byte("`/"),
		'z': []byte("2"),
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := bytes.ToLower(scanner.Bytes())
	var ans []byte
	for _, c := range in {
		if val, ok := m[c]; ok {
			ans = append(ans, val...)
		} else {
			ans = append(ans, c)
		}
	}
	fmt.Print(string(ans))
}