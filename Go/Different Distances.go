package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var x1, x2, y1, y2, p float64
		if scanner.Text() == "0" {
			break
		} else {
			fmt.Sscanf(scanner.Text(), "%f %f %f %f %f", &x1, &y1, &x2, &y2, &p)
			fmt.Println(math.Pow(math.Pow(math.Abs(x1-x2), p)+(math.Pow(math.Abs(y1-y2), p)), 1/p))
		}
	}
}
