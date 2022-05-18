package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

func swapPtr(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func main() {
	a := 1
	b := 2
	swap(a, b)
	fmt.Println(a, b)
	swapPtr(&a, &b)
	fmt.Println(a, b)
}
