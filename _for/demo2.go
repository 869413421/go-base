package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for {
		i++
		if i > 100 {
			break
		}
		sum += i
	}

	fmt.Println(sum) // 输出5050
}
