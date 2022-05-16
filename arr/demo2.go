package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Printf("Element %d is %d \n", i, a[i])
	}
}
