package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	for i := 0; i < 256; i++ {
		if A^i == B {
			fmt.Println(i)
			break
		}
	}
}
