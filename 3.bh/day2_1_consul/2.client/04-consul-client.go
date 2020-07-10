package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go_study3/3.bh/day2_1_consul/pb"
	"google.golang.org/grpc"
	"strconv"
)

func main()  {
	// 从consul 服务上获取一个健康的 服务
	// 1. 初始化 consul 的配置
	consulConfig := api.DefaultConfig()

	// 2. 获取 consul
	consulServer, err := api.NewClient(consulConfig)

	// 3. 从consul上 获取 健康的服务 -- 得一个切片.
	services, _, err := consulServer.Health().Service("grpc AND consul",
		"grpc", true, nil)

	// 将获取到的 ip 和port 拼接成一个完整的字符串.
	target := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	//////////////////// 以下是 grpc 客户端//////////////////////

	// 1. 链接 grpc服务
	//grpcConn, _ := grpc.Dial("127.0.0.1:8004", grpc.WithInsecure())
	grpcConn, _ := grpc.Dial(target, grpc.WithInsecure())

	// 2. 初始化客户端对象
	grpcClient := pb.NewHelloClient(grpcConn)

	// 3. 远程调用
	var person pb.Person
	person.Name = "Andy"
	person.Age = 18

	p, err := grpcClient.SayHello(context.TODO(), &person)

	fmt.Println(p, err)
}
