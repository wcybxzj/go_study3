package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go_study3/3.bh/day2_1_consul/pb"
	"google.golang.org/grpc"
	"net"
)

// 定义类
type Children struct {
}

// 绑定类对象
func (this *Children) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = p.Name + "grpc consul test!"
	return p, nil
}

func main() {
	// 把 grpc 服务, 注册到 consul 上.

	// 1. 初始化 consul 的配置
	consulConfig := api.DefaultConfig()

	// 2. 获取 consul --- 对应命令 consul agent -dev 开启的consul
	consulServer, err := api.NewClient(consulConfig)

	// 3. 整合即将注册到consul 上服务的配置信息 -- 初始结构体
	regSrv := api.AgentServiceRegistration{
		ID:      "bj40",
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc AND consul",
		Address: "127.0.0.1",
		Port:    8004,
		Check: &api.AgentServiceCheck{
			Name:     "consul test",
			TCP:      "127.0.0.1:8004",
			Timeout:  "3s",
			Interval: "5s",   // 不能省略 ",",否则 } 续行写.
		},  // 不能省略 ",",否则 } 续行写.
	}

	// 4. 将 grpc 服务, 注册到 consuL 上.
	consulServer.Agent().ServiceRegister(&regSrv)

	////////////////// 以下是 grpc 远程服务端代码 ////////////////

	// 1. 初始 grpc 对象
	grpcServer := grpc.NewServer()

	// 2. 注册服务
	pb.RegisterHelloServer(grpcServer, new(Children))

	// 3. 设置监听 , 指定 IP + port
	listener, err := net.Listen("tcp", "127.0.0.1:8004")
	if err != nil {
		fmt.Println("net.Listen error :", err)
		return
	}
	defer listener.Close()

	// 4. 启动服务
	grpcServer.Serve(listener)
}
