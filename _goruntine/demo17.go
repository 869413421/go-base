package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

func main() {
	pool = &sync.Pool{
		New: func() interface{} {
			return "hello,world!"
		},
	}

	value := "hello,test!"
	pool.Put(value)
	fmt.Println(pool.Get().(string))
	fmt.Println(pool.Get().(string))
	fmt.Println(pool.Get().(string))
}
