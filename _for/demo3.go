package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}

	for k, v := range a {
		fmt.Println(k, v)
	}
}
