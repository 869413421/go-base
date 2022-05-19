package main

import "fmt"

func callBack(f func()) {
	f()
}

func main() {
	var j = 1
	f := func() {
		var i = 4
		i++
		j += 1
		fmt.Printf("i, j: %d, %d\n", i, j)
	}

	f()
	j += 2
	f()
	callBack(f)
	fmt.Println(j)
}
