package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	if len(S) == 1 {
		fmt.Println(S)
		fmt.Println(S)
		return
	}
	n := len(S)
	ans_max := S
	ans_min := S
	for i := 0; i < 1100; i++ {
		S = S[1:n] + S[0:1]
		if S < ans_min {
			ans_min = S
		}
		if ans_max < S {
			ans_max = S
		}
	}
	fmt.Println(ans_min)
	fmt.Println(ans_max)
}
