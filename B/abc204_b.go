package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var ans int = 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a > 10 {
			ans += a - 10
		}
	}
	fmt.Println(ans)
}
