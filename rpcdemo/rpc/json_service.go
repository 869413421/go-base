package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpcdemo/rpc/utils"
)

type ServiceHandler struct {
}

func (*ServiceHandler) GetName(id int, item *utils.Item) error {
	log.Printf("receive GetName call, id: %d", id)
	item.Name = "xiaoming"
	item.ID = 1
	return nil
}

func (*ServiceHandler) SaveName(item utils.Item, response *utils.Response) error {
	log.Printf("receive SaveName call, Item: %v", item)
	response.OK = true
	response.ID = item.ID
	response.Message = "成功"
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
	serverHandler := &ServiceHandler{}
	err = server.Register(serverHandler)
	if err != nil {
		log.Fatalf("register error:%v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accept client error :%v", err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
