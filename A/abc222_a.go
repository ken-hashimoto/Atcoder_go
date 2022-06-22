package main

import "fmt"

func main() {
	var N string
	fmt.Scan(&N)
	for i := 0; ; i++ {
		if len(N) == 4 {
			break
		}
		N = "0" + N
	}
	fmt.Println(N)
}
