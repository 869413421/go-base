package main

import "fmt"

var d2 = 10

func Test() {
	var d2 = 20
	fmt.Printf("局部变量d2 is %d \n", d2)
}

func main() {
	Test()
	fmt.Printf("全局变量d2 is %d \n", d2)
}
