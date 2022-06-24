package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

var lock *sync.Mutex

func add(a, b int) {
	c := a + b
	lock.Lock()
	counter++
	lock.Unlock()
	fmt.Printf("%d: %d + %d = %d\n", counter, a, b, c)
}

func main() {
	startTime := time.Now()
	lock = &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go add(1, i)
	}

	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	endTime := time.Now()
	consume := endTime.Sub(startTime).Milliseconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
