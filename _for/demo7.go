package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
BREAK1:
	for i := 0; i < len(arr); i++ {
		for j := 0; j <= len(arr[i]); j++ {
			if j > 1 {
				break BREAK1
			}
			fmt.Println(arr[i][j])
		}
	}
}
