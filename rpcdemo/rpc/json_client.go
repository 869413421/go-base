package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	"rpcdemo/rpc/utils"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "localhost:8080", 30*time.Second)
	if err != nil {
		log.Fatalf("client start error :%v", err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	var item utils.Item
	client.Call("ServiceHandler.GetName", 1, &item)
	log.Printf("ServiceHandler.GetName返回结果：%v\n", item)

	var response utils.Response
	item = utils.Item{
		ID:   1,
		Name: "xiaohuang",
	}
	err = client.Call("ServiceHandler.SaveName", item, &response)
	if err != nil {
		return 
	}
	log.Fatalf("ServiceHandler.SaveName：%v\n", response)
}
