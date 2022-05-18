package main

import "fmt"

func main() {
	var a = 100
	var ptr *int

	ptr = &a
	fmt.Println(a)   // 变量值
	fmt.Println(ptr) // 变量值存储地址
}
