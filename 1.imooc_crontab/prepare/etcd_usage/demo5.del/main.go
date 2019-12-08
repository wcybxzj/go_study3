package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//删除
/*
先运行下有数据再删
./etcdctl put  /cron/jobs/job2 2222
OK

./etcdctl put  /cron/jobs/job1 1111
OK

./etcdctl get --from-key=true /cron/jobs
/cron/jobs/job1
1111
/cron/jobs/job2
2222
*/
func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		delResp *clientv3.DeleteResponse
		//kvPair  *mvccpb.KeyValue
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

	// 删除KV
	//  clientv3.WithFromKey(): 从这个key开始

	delResp, err = kv.Delete(context.TODO(), "/cron/jobs",
							clientv3.WithFromKey(), clientv3.WithPrevKV())

	if err != nil {
		fmt.Println(err)
		return
	}else {
		fmt.Println("del ok")
		fmt.Println("delResp len():", len(delResp.PrevKvs))
	}

	// 被删除之前的value是什么
	if len(delResp.PrevKvs) == 0 {
		return
	}

	for _, kvPair := range delResp.PrevKvs {
		fmt.Println("删除了:", string(kvPair.Key), string(kvPair.Value))
	}

}

