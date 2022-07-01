package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpcdemo/rpc/utils"
)

type MathService struct {
}

func (*MathService) Multiply(args *utils.MathArgs, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (*MathService) Divide(args *utils.MathArgs, reply *int) error {
	if args.B == 0 {
		return errors.New("不可以除以0")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	//1.注册服务，并且指定协议为http
	math := new(MathService)
	rpc.Register(math)
	rpc.HandleHTTP()

	//2.监听端口
	listen, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatal(err, "服务启动失败")
	}
	fmt.Println("服务启动成功")
	//3.启动服务
	err = http.Serve(listen, nil)
	if err != nil {
		log.Fatal(err, "服务启动失败")
	}
	fmt.Println("服务启动成功")

}
