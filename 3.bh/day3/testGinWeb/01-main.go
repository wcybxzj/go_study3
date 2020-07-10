package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"

	testSrv "go_study3/3.bh/day3/testGinWeb/proto/testSrv" // 最前端 的 testSrv 是别名

	"context"
	"fmt"
	"github.com/micro/go-plugins/registry/consul"
)

// 创建 实名函数, 供 GET请求 回调
func CallRemote(ctx *gin.Context)  {

	// 初始化consuL服务发现
	consulReg := consul.NewRegistry()

	// 初始化 micro对象, 指定 consul 为服务发现
	service := micro.NewService(micro.Registry(consulReg))

	// 1. 初始化 访问testSrv 微服务的 "客户端"
	microClient := testSrv.NewTestSrvService("go.micro.srv.testSrv", service.Client())

	// 2. 调用 testSrv 提供的远程的函数
	resp, err := microClient.Call(context.TODO(), &testSrv.Request{Name:"xiaowang"})

	// 3. 将 远程调用返回的结果, 显示回给 浏览器
	ctx.Writer.WriteString(resp.Msg)

	fmt.Println(resp, err)
}

func main()  {

	// 1. 初始化路由 --- [初始化web引擎]
	router := gin.Default()

	// 2. 路由匹配
	router.GET("/", CallRemote)

	// 3. 运行
	// router.Run("192.168.6.108:8091")
	router.Run(":8091")
}
