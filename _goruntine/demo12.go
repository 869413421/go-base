package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

func addSum(a, b int) {
	defer wg.Done()
	c := a + b
	fmt.Printf("%d + %d = %d \n", a, b, c)
}

func main() {
	wg = new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go addSum(1, i)
	}

	wg.Wait()
}
