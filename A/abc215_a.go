package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	if S == "Hello,World!" {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}
