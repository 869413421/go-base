package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	keys := make([]string, 3)
	for k := range map1 {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, map1[v])
	}

	invMap := make(map[int]string, 3)
	for k, v := range map1 {
		invMap[v] = k
	}

	keys2 := make([]int, 3)
	for k, _ := range invMap {
		keys2 = append(keys2, k)
	}
	sort.Ints(keys2)
	for _, v := range keys2 {
		fmt.Println(v, invMap[v])
	}
}
