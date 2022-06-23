package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	edge_cnt := make(map[int]int)
	var a, b int
	for i := 0; i < N-1; i++ {
		fmt.Scan(&a, &b)
		if _, ok := edge_cnt[a]; ok {
			edge_cnt[a] += 1
		} else {
			edge_cnt[a] = 1
		}
		if _, ok := edge_cnt[b]; ok {
			edge_cnt[b] += 1
		} else {
			edge_cnt[b] = 1
		}
	}
	for _, val := range edge_cnt {
		if val == N-1 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
