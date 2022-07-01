package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpcdemo/rpc/utils"
)

func main() {
	serviceAddress := "localhost:8989"
	client, err := rpc.DialHTTP("tcp", serviceAddress)
	if err != nil {
		log.Fatal(err, "连接服务端失败")
	}
	fmt.Println("客户端启动成功")
	//1.同步调用
	args := utils.MathArgs{A: 10, B: 10}
	var reply int
	err = client.Call("MathService.Multiply", &args, &reply)
	if err != nil {
		log.Fatal(err, "rpc request MathService.Multiply error")
	}
	fmt.Printf("同步：%d*%d=%d\n", args.A, args.B, reply)

	divideCall := client.Go("MathService.Divide", &args, &reply, nil)

	for{
		select {
		case <-divideCall.Done:
			fmt.Printf("异步：%d/%d=%d\n", args.A, args.B, reply)
			return
		}
	}

}
