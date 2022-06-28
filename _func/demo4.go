package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 2))

	c := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(c)
}
