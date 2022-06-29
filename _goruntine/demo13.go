package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func LoadConfig() {
	fmt.Println("Start !!!")
	once.Do(func() {
		fmt.Println("do Something !!!")
	})
	fmt.Println("End !!!")
}

func main() {
	LoadConfig()
	LoadConfig()
}
