package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 1; ; i++ {
		if (i-1)*100+1 <= N && N <= i*100 {
			fmt.Println(i)
			break
		}
	}
}
