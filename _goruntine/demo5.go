package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.创建生产管道
	ch := make(chan int, 10)

	// 2.创建超时管道
	timeout := make(chan bool)

	// 3.开启超时携程
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	// 4.使用select结束阻塞
	select {
	case <-ch:
		fmt.Println("接收到数据")
	case <-timeout:
		fmt.Println("执行超时！")
	}
}
