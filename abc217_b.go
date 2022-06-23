package main

import "fmt"

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)
	arr := [...]string{"ABC", "AGC", "ARC", "AHC"}
	for i := range arr {
		if arr[i] == s1 || arr[i] == s2 || arr[i] == s3 {
			arr[i] = ""
		}
	}
	for i := range arr {
		if arr[i] != "" {
			fmt.Println(arr[i])
			return
		}
	}
}
