package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

var num1 int32 = 0

func testAdd(deferFunc func()) {
	defer deferFunc()
	atomic.AddInt32(&num1, 1)
	fmt.Println(num1)
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 10; i < 100; i++ {
		go testAdd(func() {
			fmt.Println("defer")
			fmt.Println(atomic.LoadInt32(&num1))
			if atomic.LoadInt32(&num1) >= 10 {
				cancelFunc()
			}
		})
	}
	<-ctx.Done()
	fmt.Println("完成")
}
