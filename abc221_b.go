package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)
	if S == T {
		fmt.Println("Yes")
		return
	}
	l := len(S)
	var i int = 0
	for i = 0; i < l-1; i++ {
		if S[i] != T[i] && S[i+1] != T[i+1] {
			if S[i+1] == T[i] && T[i+1] == S[i] {
				fmt.Println("Yes")
				return
			} else {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("No")
}
