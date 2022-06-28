package main

import "fmt"

type String1 interface {
	func1(s string)
	func2(s string)
}

type String2 interface {
	func1(s string)
}

type String string

func (s String) func1(str string) {
	fmt.Println(str)
}

func (s String) func2(str string) {

}

func main() {
	var s String = "test"
	var s1 String2 = s
	if s2, ok := s1.(String1); ok {
		s2.func1("test2")
	}
}
