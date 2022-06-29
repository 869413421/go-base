package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("发送%d \n", i)
		}
		close(ch)
	}()

	for {
		num, ok := <-ch
		if ok {
			fmt.Printf("接收%d \n", num)
		} else {
			fmt.Println("结束")
			break
		}
	}
}
