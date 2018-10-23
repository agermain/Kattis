package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := strings.Split(scanner.Text(), " ")
	num, _ := strconv.Atoi(n[0])
	s := "1"
	if len(n) > 1 {
		for _, v := range n[1] {
			if v == 82 {
				s += "1"
			} else {
				s += "0"
			}
		}
	}
	sn, _ := strconv.ParseInt(s, 2, 64)
	fmt.Println(int64((math.Pow(2, float64(num+1))) - float64(sn)))
}
