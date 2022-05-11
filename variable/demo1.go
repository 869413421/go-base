package main

import "fmt"

// 定义一个
var i int

// 定义多个同类型
var i1, i2 int

// 分组定义
var (
	a, a1 int
	b     string
)

func main() {
	fmt.Printf("i is %d \n", i)
	fmt.Printf("i1 is %d \n", i1)
	fmt.Printf("i2 is %d \n", i2)
	fmt.Printf("a is %d \n", a)
	fmt.Printf("a1 is %d \n", a1)
	fmt.Printf("b is %s \n", b)
}
