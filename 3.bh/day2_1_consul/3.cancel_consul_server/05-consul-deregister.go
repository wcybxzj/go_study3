package main

import (
	"github.com/hashicorp/consul/api"
	"fmt"
)

//注销服务
func main()  {

	// 1. 初始化 consul 的配置
	consulConfig := api.DefaultConfig()

	// 2. 获取 consul --- 对应命令 consul agent -dev 开启的consul
	consulServer, err := api.NewClient(consulConfig)
	fmt.Println("err:", err)

	// 3. 注销服务
	err = consulServer.Agent().ServiceDeregister("bj40")

	fmt.Println("err:", err)
}

