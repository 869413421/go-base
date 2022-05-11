package main

import "fmt"

func main() {
	var d int
	fmt.Printf("d is %d \n", d) // 输出0
	d = 100
	fmt.Printf("d is %d \n", d) // 输出100
	d = 200
	fmt.Printf("d is %d \n", d) // 输出200

	d1 := 1
	d2 := 2
	fmt.Printf("d1 is %d \n", d1)
	fmt.Printf("d2 is %d \n", d2)
	d1, d2 = d2, d1
	fmt.Printf("d1 is %d \n", d1)
	fmt.Printf("d2 is %d \n", d2)
}
