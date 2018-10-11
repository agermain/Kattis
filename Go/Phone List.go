package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := toInt(scanner.Bytes())
	for i:=0; i<n; i++ {
		b := false
		scanner.Scan()
		m := toInt(scanner.Bytes())
		numbers := make([]string, m)
		for j:=0; j<m; j++ {
			scanner.Scan()
			numbers[j] = scanner.Text()
		}
		sort.Strings(numbers)
		for j:=1; j<len(numbers); j++ {
			if strings.HasPrefix(numbers[j], numbers[j-1]){
				b = true
				break
			}
		}
		if b {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}