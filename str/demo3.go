package main

import "fmt"

func main() {
	str := "hello,世界"
	for i, ch := range str {
		fmt.Println(i, string(ch)) // ch 的类型为 rune
	}
}
