package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}

	// 复制slice1到slice2
	//copy(slice2, slice1)
	//fmt.Println(slice2) // 输出1,2,3

	// 复制slice2到slice1
	copy(slice1, slice2)
	fmt.Println(slice1) // 输出6,7,8,4,5
}
