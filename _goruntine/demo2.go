package main

import (
	"fmt"
	"time"
)

func channelAdd(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d+%d = %d \n", a, b, c)
	ch <- c
}

func main() {
	start := time.Now()
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go channelAdd(1, i, chs[i])
	}

	for _, ch := range chs {
		<- ch
	}


	end := time.Now()
	consume := string(end.Sub(start).String())
	fmt.Printf("coumse %s s", consume)

}
