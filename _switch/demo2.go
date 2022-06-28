package main

import "fmt"

func main() {
	num := 100
	switch {
	case num > 90:
		fmt.Println("a")
		fallthrough
	case num > 80:
		fmt.Println("b")
	case num > 70 && num == 100:
		fmt.Println("c")
	}
	//输出a,b
}
