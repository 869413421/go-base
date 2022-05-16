package main

import "fmt"

func main() {
	var a float64= 0.7
	var b float64 = 0.1
	c := a + b
	fmt.Printf("a+b=%v \n", c) // 输出0.7999999999999999
	fmt.Printf("0.8 == 0.1+0.7 ? %v", (0.8 == a+b)) // 输出false
}
