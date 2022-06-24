package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	chs := []chan int{
		make(chan int, 10),
		make(chan int, 10),
		make(chan int, 10),
	}

	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(1 * time.Second)
		index := rand.Intn(3)
		fmt.Println(index)
		chs[2] <- 3
		select {
		case <-chs[0]:
			fmt.Println("通道一")
		case <-chs[1]:
			fmt.Println("通道二")
		case num := <-chs[2]:
			fmt.Println("通道三", num)
		default:
			fmt.Println("默认分支")
		}
	}
}
