package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var k int = 0
	var now int = 1
	for k < 10000 {
		now *= 2
		if now > n {
			fmt.Println(k)
			break
		}
		k++
	}
}
