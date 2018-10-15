package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

func Round(x, unit float64) float64 {
    return math.Round(x/unit) * unit
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    var n, m float64
    fmt.Sscanf(scanner.Text(), "%f %f", &n, &m)
    fmt.Println(int(Round(n,math.Pow(10, m))))
}