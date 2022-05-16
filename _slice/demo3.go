package main

import "fmt"

func main() {
	slice1 := make([]int, 4, 10)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := append(slice1, 1, 2, 3)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

}
