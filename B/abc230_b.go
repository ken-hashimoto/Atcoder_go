package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	T := "oxxoxxoxxoxxoxxoxxoxxoxxoxxoxxoxxoxx"
	for i := 0; i < 10; i++ {
		if S == T[i:len(S)+i] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
