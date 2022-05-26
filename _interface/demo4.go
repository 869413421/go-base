package main

import "fmt"

type Number1 interface {
	Equal(i int) bool
	Less(i int) bool
	More(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	Less(i int) bool
	More(i int) bool
	Add(i int)
}

type Number int

func (n Number) Equal(i int) bool {
	return int(n) == i
}
func (n Number) Less(i int) bool {
	return int(n) < i
}
func (n Number) More(i int) bool {
	return int(n) > i
}

func main() {
	var num Number = 1
	var num2 Number1
	num2 = num
	//var num3 Number2 = num2
	//num2 = num3
	fmt.Println(num2)
}
