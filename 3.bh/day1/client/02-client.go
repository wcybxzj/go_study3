package main

import (
	"fmt"
	"go_study3/3.bh/day1/pb"

	//"net/rpc/jsonrpc"
)

/*func main()  {
	// 1. 使用rpc链接服务端
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8000")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 调用远端函数, [对象.函数名(服务名.方法名, 传入, 传出)]
	var reply string
	err = conn.Call("hello.HelloWorld", "李白", &reply)
	if err != nil {
		fmt.Println("Call err:", err)
		return
	}

	fmt.Println(reply)
}*/

// 使用封装后的 2.client, 调用远程函数
//go run 01-1.server.go 03-design.go
func main()  {
	myClient := pb.InitClient("127.0.0.1:8000")

	var reply string

	err := myClient.HelloWorld("杜甫", &reply)

	fmt.Println(reply, err)
}
