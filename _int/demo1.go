package main

import "fmt"

func main() {

	var num1 int32 = 10
	num2 := 8                       // 自动推导为int
	var num3 int = int(num1) + num2 // 需要对类型进行转换才可以运算
	fmt.Println(num3)

	var a int = 8
	var b int32 = 10

	c := a + int(b)
	c++
	fmt.Println(c)
	c += c
	fmt.Println(c)
	c -= c
	fmt.Println(c)
}
