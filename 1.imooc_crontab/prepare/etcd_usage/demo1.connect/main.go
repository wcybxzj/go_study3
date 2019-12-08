package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//简单连接
func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)

	// 客户端配置
	config = clientv3.Config{
		//		Endpoints: []string{"36.111.184.221:2379"},
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	client = client
}
