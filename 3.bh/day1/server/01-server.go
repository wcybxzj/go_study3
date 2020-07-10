package main

import (
	"fmt"
	"go_study3/3.bh/day1/pb"
	"net"
	"net/rpc/jsonrpc"
)

//go run 01-1.server.go 03-design.go
func main() {
	// 1. 注册rpc服务,绑定对象
	/*
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("注册rpc失败:", err)
		return
	}*/

	pb.RigsterService(new(pb.World))

	// 2. 启动监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听...")

	for {


	// 3. 建立链接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("链接成功...")

	// 4. 绑定rpc 服务
	jsonrpc.ServeConn(conn)

	}
}
