package main

import "fmt"

func main() {
	var a interface{} = 1
	var b interface{} = "sss"
	var c interface{} = 1.11
	var d interface{} = true
	var e interface{} = []int{1, 2, 3}
	var f interface{} = [2]int{1, 2}
	var g interface{} = struct {
		id int
	}{id: 1}
}

func test(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
