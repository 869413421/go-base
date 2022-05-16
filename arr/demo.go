package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	fmt.Println(a)
	b := [5]int{1: 1, 3: 5}
	fmt.Println(b)
	fmt.Println(b[3])
	fmt.Println(b[len(b)-1])
}
