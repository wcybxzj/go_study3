package main

import (
	"context"
	"fmt"
	"go_study3/3.bh/day2/pb"
	"google.golang.org/grpc"
)

func main()  {
	//1.  连接 grpc 服务
	grpcConn, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())  // 参2: 安全的. 固定写法.
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer grpcConn.Close()

	//2.  初始化 grpc 客户端
	grpcClient := pb.NewSayNameClient(grpcConn)

	var teacher pb.Teacher

	teacher.Name = "itcast"
	teacher.Age = 100

	//3.  调用远程函数
	t, err := grpcClient.SayHello(context.TODO(), &teacher)  // 参1: context.TODO() 固定写法.

	fmt.Println(t, err)
}
