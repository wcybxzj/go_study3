package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"testSrv/handler"

	"github.com/micro/go-plugins/registry/consul"
	testSrv "testSrv/proto/testSrv" // 最前端的 testSrv 是后面目录的别名
)

func main() {
	// 初始化 服务发现  --- 使用 默认属性 consul agent -dev 的属性
	consulReg := consul.NewRegistry()

	// New Service --- 创建 micro 服务
	service := micro.NewService(
		micro.Name("go.micro.srv.testSrv"),
		micro.Registry(consulReg),  // 使用 consul 做服务发现
		micro.Version("latest"),
	)

	// Register Handler
	testSrv.RegisterTestSrvHandler(service.Server(), new(handler.TestSrv))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
