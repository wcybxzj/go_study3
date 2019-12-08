package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//op+do操作取代 get/put/delete
//目的:为学习分布式锁做准备

//第一次执行
//写入Revision: 109
//CreateRevision: 109  -------------
//ModRevision: 109
//数据value: 123123123

//第二次执行
//写入Revision: 110
//CreateRevision: 109 --------------
//ModRevision: 110
//数据value: 123123123

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
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

	kv = clientv3.NewKV(client)

	// 创建Op: operation
	putOp = clientv3.OpPut("/cron/jobs/job8", "123123123")

	// 执行OP
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	// kv.Do(op)
	// kv.Put
	// kv.Get
	// kv.Delete

	fmt.Println("1.写入Revision(被修改就会变化):", opResp.Put().Header.Revision)

	// 创建Op
	getOp = clientv3.OpGet("/cron/jobs/job8")

	// 执行OP
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	// 打印
	fmt.Println("2.CreateRevision(被创建的版本):", opResp.Get().Kvs[0].CreateRevision) // create rev == mod rev
	fmt.Println("3.ModRevision（被修改的版本）:", opResp.Get().Kvs[0].ModRevision) // create rev == mod rev

	fmt.Println("数据value:", string(opResp.Get().Kvs[0].Value))
}
