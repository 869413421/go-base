package main

import "fmt"

type Integer int

func (a Integer) Equal(b Integer) bool {
	return a == b
}

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a Integer) More(b Integer) bool {
	return a > b
}

func (a *Integer) Increase(i Integer) {
	*a += i
}

func (a *Integer) Decrease(i Integer) {
	*a -= i
}

func main() {
	var a Integer = 2
	if a.Equal(2) {
		fmt.Println("a is equal to 2")
	} else {
		fmt.Println("a not equal to 2")
	}

	a.Increase(2)
	fmt.Println(a)
	a.Decrease(2)
	fmt.Println(a)
}
