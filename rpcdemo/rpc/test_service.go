package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpcdemo/rpc/utils"
)

type MService struct {
}

func (*MService) Add(data interface{}, response *utils.Response) error {
	log.Printf("receive Add call, Item: %v", data)
	response.OK = true
	response.ID = 0
	dataJson, err := json.Marshal(data)
	if err != nil {
		response.Message = ""
	} else {
		response.Message = string(dataJson)
	}
	return nil
}

func main() {
	//1.初始化rpc客户端
	server := rpc.NewServer()

	//2.监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("监听端口失败：%v", err)
	}
	defer listener.Close()

	log.Println("Start listen on port 8080")

	//3.初始化服务处理器
	serverHandler := &MService{}
	err = server.Register(serverHandler)
	if err != nil {
		log.Fatalf("register error:%v", err)
	}

	for {
		conn, err := listener.Accept()
		fmt.Println(&conn)
		if err != nil {
			log.Fatalf("accept client error :%v", err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
