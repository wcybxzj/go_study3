package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//GET值
func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		getResp *clientv3.GetResponse
	)

	config = clientv3.Config{
		//		Endpoints: []string{"36.111.184.221:2379"}, // 集群列表
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)


	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1" /*clientv3.WithCountOnly()*/); err != nil {
		fmt.Println(err)
	} else {
		//输出:[key:"/cron/jobs/job1" create_revision:62 mod_revision:67 version:4 value:"bye" ] 1
		//create_revision:创建版本 是对于整个etcd的版本号
		//mod_revision:修改版本
		//version:修改版本-创建版本
		//value:之前的值
		//getResp.Count:返回的数量
		//getResp.More:用来分页

		fmt.Println(getResp.Kvs,getResp.Header.Revision, getResp.Count)
	}
}
