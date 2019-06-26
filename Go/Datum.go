package main

import (
	"fmt"
	"time"
)

func main() {
	var d int
	var m time.Month
	fmt.Scanf("%d %d", &d, &m)
	t := time.Date(2009, m, d, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(t.Weekday())
}
