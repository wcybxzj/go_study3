package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//PUT值
func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		putResp *clientv3.PutResponse
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

	//获取之前这个key的值 clientv3.WithPrevKV()
	val := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("本次设置的val:" + val)
	putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", val, clientv3.WithPrevKV())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("当前的版本:", putResp.Header.Revision)
		//获取之前这个key的值 如果被覆盖
		if putResp.PrevKv != nil {
			fmt.Println("之前的值是:", string(putResp.PrevKv.Value))
		}
		fmt.Println(putResp)
	}
}
