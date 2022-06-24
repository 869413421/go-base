package main

import (
	"fmt"
	"time"
)

func test(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	start := time.Now()
	ch := make(chan int, 20)
	go test(ch)

	for v := range ch {
		fmt.Println(v)
	}
	end := time.Now()
	consume := end.Sub(start).String()
	fmt.Printf("consume %s s", consume)
}
