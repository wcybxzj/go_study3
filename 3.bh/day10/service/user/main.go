package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"bj40ihome/service/user/handler"

	user "bj40ihome/service/user/proto/user"
	"github.com/micro/go-micro/registry/consul"
	"bj40ihome/service/user/model"
)

func main() {
	// 初始化 MySQL 连接池
	model.InitDb()

	// 初始化 redis 连接池
	model.InitRedis()

	// 初始化 consul配置
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Address("192.168.6.108:12342"),   // 指定微服务端口号
		micro.Name("go.micro.srv.user"),
		micro.Registry(consulReg),			// 指定使用的 服务发现
		micro.Version("latest"),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
