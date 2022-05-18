package main

import (
	"fmt"
	"strconv"
)

func main() {
	var map1 map[string]int
	fmt.Println(map1)
	map2 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(map2)
	key := "two"
	value, ok := map2[key]
	if ok {
		fmt.Printf("map has key:" + key + " values is:" + strconv.Itoa(value))
	} else {
		fmt.Printf("map no has key")
	}
}
