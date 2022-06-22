package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	x, _ := strconv.Atoi(s[0:1])
	y, _ := strconv.Atoi(s[2:])
	fmt.Println(x * y)
}
