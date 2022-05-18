package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["test"] = 1
	map1["test2"] = 2
	fmt.Println(map1)
	delete(map1, "test")
	fmt.Println(map1)
}
