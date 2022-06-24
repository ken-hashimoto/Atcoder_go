package main

import (
	"fmt"
	"sort"
)

type player struct {
	index, score int
}

func main() {
	var n int
	fmt.Scan(&n)
	players := make([]player, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		players[i].index = i + 1
		players[i].score = a
	}
	sort.Slice(players, func(i, j int) bool { return players[i].score < players[j].score })
	fmt.Println(players[n-2].index)
}
