package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	vote := make(map[string]int)
	var S string
	for i := 0; i < N; i++ {
		fmt.Scan(&S)
		if _, ok := vote[S]; ok {
			vote[S] += 1
		} else {
			vote[S] = 1
		}
	}
	var ans string
	var max_vote int
	for key, val := range vote {
		if max_vote < val {
			ans = key
			max_vote = val
		}
	}
	fmt.Println(ans)
}
