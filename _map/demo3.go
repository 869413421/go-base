package main

import "fmt"

func main() {
	map1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	invMap := make(map[int]string, 3)

	for k, v := range map1 {
		invMap[v] = k
	}

	for k, v := range invMap {
		fmt.Println(k, v)
	}
}
