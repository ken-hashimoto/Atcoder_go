package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	ans := 1
	for i := 0; i < A-B; i++ {
		ans *= 32
	}
	fmt.Println(ans)
}
