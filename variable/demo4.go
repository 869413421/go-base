package main

import "fmt"

func GetName() (name, nickName string) {
	name = "test"
	nickName = "测试"
	return name, nickName
}

func main() {
	_, nickName := GetName()
	fmt.Printf("nickName is %s \n", nickName)
}
