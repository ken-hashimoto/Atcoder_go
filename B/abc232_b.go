package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	diff := int(t[0]) - int(s[0])
	if diff < 0 {
		diff += 26
	}
	for i := 1; i < len(s); i++ {
		tmp := int(t[i]) - int(s[i])
		if tmp < 0 {
			tmp += 26
		}
		if tmp != diff {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
