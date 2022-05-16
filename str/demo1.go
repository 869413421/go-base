package main

import "fmt"

func main() {
	str := "hello,world"
	str1 := str[:5]  //取下标五之前的数据
	str2 := str[7:]  // 取下标7后的数据
	str3 := str[1:4] // 去下标1到4前的数据

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
}
