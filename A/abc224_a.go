package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	l := len(S)
	if S[l-2:] == "er" {
		fmt.Println("er")
	} else if l >= 3 && S[l-3:] == "ist" {
		fmt.Println("ist")
	}
}
