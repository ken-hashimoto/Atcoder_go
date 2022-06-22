package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	var s1, s2, s3 = S[0:1], S[1:2], S[2:]
	distinct := 1
	if s1 != s2 || s1 != s3 || s2 != s3 {
		distinct++
	}
	if s1 != s2 && s1 != s3 && s2 != s3 {
		distinct++
	}
	switch distinct {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(3)
	case 3:
		fmt.Println(6)
	}
}
