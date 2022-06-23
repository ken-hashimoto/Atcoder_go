package main

import "fmt"

func main() {
	var H, W int
	fmt.Scan(&H, &W)
	arr := make([][]int, H)
	for i := range arr {
		arr[i] = make([]int, W)
		for j := range arr[i] {
			fmt.Scan(&arr[i][j])
		}
	}
	for i1 := 0; i1 < H; i1++ {
		for i2 := i1 + 1; i2 < H; i2++ {
			for j1 := 0; j1 < W; j1++ {
				for j2 := j1 + 1; j2 < W; j2++ {
					if arr[i1][j1]+arr[i2][j2] > arr[i2][j1]+arr[i1][j2] {
						fmt.Println("No")
						return
					}
				}
			}
		}
	}
	fmt.Println("Yes")
}
