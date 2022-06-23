package main

import "fmt"

func main() {
	var N, P int
	var ans int
	fmt.Scan(&N, &P)
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		if a < P {
			ans++
		}
	}
	fmt.Println(ans)
}
