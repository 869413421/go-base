package main

import (
	"context"
	"fmt"
	"time"
)

func go1(ctx context.Context) {
	ch := make(chan bool)
	go func() {
		time.Sleep(20 * time.Microsecond)
		ch <- true
	}()
	select {
	case <-ch:
		fmt.Println("正常结束")
		return
	case <-ctx.Done():
		fmt.Println("go1 手动结束或超时")
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Microsecond)
	go go1(ctx)
	i := 0
	for {
		i++
		time.Sleep(10 * time.Microsecond)
		if i > 10000 {
			fmt.Println("手动结束")
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Second)
}
