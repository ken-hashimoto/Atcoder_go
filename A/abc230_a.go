package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var N string
	fmt.Scan(&N)
	var l = utf8.RuneCountInString(N)
	var n, _ = strconv.Atoi(N)
	if n >= 42 {
		n++
	}
	N = strconv.Itoa(n)
	switch l {
	case 1:
		fmt.Println("AGC00" + N)
	case 2:
		fmt.Println("AGC0" + N)
	}
}
