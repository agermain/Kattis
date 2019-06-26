package main

import "fmt"

func main() {

	for {
		var a int
		fmt.Scanln(&a)
		if a == -1 {
			break
		}
		var finalT, distance = 0, 0
		for i := 0; i < a; i++ {
			var speed, time int
			fmt.Scanf("%d %d", &speed, &time)
			time -= finalT
			finalT += time
			distance += time * speed
		}
		fmt.Printf("%d miles\n", distance)
	}
}
