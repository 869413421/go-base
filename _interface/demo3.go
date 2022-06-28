package main

import "fmt"

type Integer int

func (i*Integer) Add(a, b Integer) Integer {
	return a + b
}

func (i Integer) Multiply(b Integer) Integer {
	return i * b
}

type Math interface {
	Add(a, b Integer) Integer
	Multiply(i Integer) Integer
}

func main() {
	var a Integer = 2
	var m Math
	m = &a
	fmt.Println(m.Add(2, 1))
	fmt.Println(m.Multiply(2))
	fmt.Println(a)
}
