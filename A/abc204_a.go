package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x == y {
		fmt.Println(x)
	} else {
		if x > y {
			x, y = y, x
		}
		if x == 0 && y == 1 {
			fmt.Println(2)
		} else if x == 1 && y == 2 {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
