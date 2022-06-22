package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)
	if S < T {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
