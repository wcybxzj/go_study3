package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"bj40ihome/service/getCaptcha/handler"

	getCaptcha "bj40ihome/service/getCaptcha/proto/getCaptcha"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	// 初始化 consul
	consulReg := consul.NewRegistry()		// 默认属性 consul agent -dev

	// New Service
	service := micro.NewService(
		micro.Address("192.168.6.108:12341"),
		micro.Name("go.micro.srv.getCaptcha"),
		micro.Registry(consulReg),    // 指定 使用 consul服务发现
		micro.Version("latest"),
	)

	// Register Handler
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
