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

	var response utils.Response
	data := utils.Item{
		ID:   1,
		Name: "小猪",
	}
	err = client.Call("MService.Add", data, &response)
	if err != nil {
		return
	}
	log.Fatalf("MService.Add：%v\n", response)
}
