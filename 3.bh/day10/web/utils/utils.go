package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
)

// 初始化微服务对象
func InitMicro() micro.Service {
	// 创建 consul 服务发现
	consulReg := consul.NewRegistry()
	// 获取 微服务对象
	return micro.NewService(micro.Registry(consulReg))
}
