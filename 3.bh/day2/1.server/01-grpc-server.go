package main

import (
	"google.golang.org/grpc"
	"go_study3/3.bh/day2/pb"
	"context"
	"net"
	"fmt"
)
// 定义类
type Children struct {
}

// 绑定类方法
func (this *Children)SayHello(ctx context.Context, t *pb.Teacher) (*pb.Teacher, error) {
	t.Name = t.Name + " is a Teacher!"
	t.Age = 28;
	return t, nil;
}

func main()  {
	// 1. 初始化一个 grpc 对象
	grpcServer := grpc.NewServer()

	//2.  注册服务
	pb.RegisterSayNameServer(grpcServer, new(Children))

	//3.  设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("启动服务...")

	//4.  启动服务 —— 使用 创建的 grpc 对象。
	grpcServer.Serve(listener)

}
