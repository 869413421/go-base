package main

import "fmt"

func main() {
	str := "hello,world!我"
	n := len(str)

	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, string(ch)) // 根据下标取字符串中的字符，类型为byte
	}
}
