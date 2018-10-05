package main

import "fmt"

func main() {
	var s string
	var c int
	fmt.Scanln(&s)
	for i:=0; i<len(s); i+=3 {
		if string(s[i]) != "P" {
			c++
		}
		if string(s[i+1]) != "E" {
			c++
		}
		if string(s[i+2]) != "R" {
			c++
		}
	}
	fmt.Println(c)
}