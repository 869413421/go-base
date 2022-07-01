package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup
var pool *sync.Pool

func initPool() {
	pool.Put("hello,init!")
	wg.Done()
}

func main() {
	wg = new(sync.WaitGroup)
	pool = &sync.Pool{
		New: func() interface{} {
			return "hello,world"
		},
	}
	wg.Add(1)
	go initPool()
	wg.Wait()

	fmt.Println(pool.Get())
}
