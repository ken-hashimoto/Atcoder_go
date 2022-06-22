package main

import "fmt"

func main() {
	var X int
	fmt.Scan(&X)
	if X > 0 && X%100 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
