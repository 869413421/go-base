package main

import "fmt"

func addFunc(a int) func(b int) int {
	return func(b int) int {
		b++
		return b
	}
}

func main() {
	f := addFunc(1)
	fmt.Println(f(1))
}
