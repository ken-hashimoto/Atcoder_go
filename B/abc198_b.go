package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("Yes")
		return
	}
	for {
		if n%10 == 0 {
			n = n / 10
		} else {
			break
		}
	}
	n_str := strconv.Itoa(n)
	l := len(n_str)
	for i := 0; i < l/2; i++ {
		if n_str[i] != n_str[l-i-1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
