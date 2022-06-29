package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter int32 = 0

func testAdd(i int32) {
	atomic.AddInt32(&counter, 1)
	//counter += 1
	fmt.Println(counter)
}

func main() {
	for i := 0; i < 100; i++ {
		go testAdd(int32(i))
		go testAdd(int32(i))
		go testAdd(int32(i))
	}

	time.Sleep(1 * time.Second)
}
