package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

//租约
//给一个key关联1个10秒的租约,10秒后租约到期key被删除
func main() {
	fmt.Println()

	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		kv clientv3.KV
	)

	config = clientv3.Config{
//		Endpoints: []string{"36.111.184.221:2379"}, // 集群列表
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 申请一个lease（租约）
	lease = clientv3.NewLease(client)

	// 申请一个10秒的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约的ID
	leaseId = leaseGrantResp.ID

	// 获得kv API子集
	kv = clientv3.NewKV(client)

	// Put一个KV, 让它与租约关联起来, 从而实现10秒后自动过期
	val := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("val:"+val)
	putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", val, clientv3.WithLease(leaseId))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入成功:", putResp.Header.Revision)

	// 定时的看一下key过期了没有
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期:", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
