package main

import "fmt"

func main() {
	JLoop:
	for i := 0; i <= 100; i++ {
		if i == 5 {
			fmt.Println("continue 跳出")
			continue
		}
		for j := 0; j <= 100; j++ {
			if i > 5 {
				break JLoop
			}
		}
	}
}
