package main

import "fmt"

func main() {
	num := 100
	switch num {
	case 90, 100:
		fmt.Println(1)
	case 80:
		fmt.Println(2)
	}

	switch {
	case num >= 90:
		fmt.Println("a")
	case num >= 90 && num < 95:
		fmt.Println("b")
	}
}
