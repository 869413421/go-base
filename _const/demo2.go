package main

import "fmt"

const (
	a1 = iota
	a2 = iota
	a3 = iota
)

const (
	b1 = iota * 2
	b2 = iota * 2
	b3 = iota * 2
)

const (
	c1 = iota * 3
	c2
	c3
)

func main() {
	fmt.Printf("a1 is %d \n", a1)
	fmt.Printf("a2 is %d \n", a2)
	fmt.Printf("a3 is %d \n", a3)
	fmt.Printf("b1 is %d \n", b1)
	fmt.Printf("b2 is %d \n", b2)
	fmt.Printf("b3 is %d \n", b3)
	fmt.Printf("c1 is %d \n", c1)
	fmt.Printf("c2 is %d \n", c2)
	fmt.Printf("c2 is %d \n", c3)
}
