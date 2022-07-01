package main

import (
	"context"
	"fmt"
	"time"
)

func go2(ctx context.Context) {
	select {
	case <-ctx.Done():
		println("携程2已结束")
		return
	}
}

func go1(ctx context.Context) {
	go go2(ctx)
	select {
	case <-ctx.Done():
		println("携程1已结束")
		return
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go go1(ctx)
	for i := 1; i < 100; i++ {
		if i > 10 {
			cancelFunc()
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Println("主携程结束")
}
