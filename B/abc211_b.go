package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1, s2, s3, s4 string
	want := []string{"2B", "3B", "H", "HR"}
	fmt.Scan(&s1, &s2, &s3, &s4)
	now := []string{s1, s2, s3, s4}
	sort.Slice(now, func(i, j int) bool { return now[i] < now[j] })
	for i := 0; i < 4; i++ {
		if now[i] != want[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
