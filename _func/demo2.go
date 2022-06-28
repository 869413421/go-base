package main

import "fmt"

func myFunc(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func main() {
	myFunc(1, 2, 3, 4, 5)
	slice1 := []int{7, 8, 9, 10}
	myFunc(slice1...)
}
