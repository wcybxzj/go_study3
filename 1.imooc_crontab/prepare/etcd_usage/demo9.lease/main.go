package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

//1.目的:实现一个乐观锁（tnx事务 + op操作来实现）
//txn事务: if else then
//op操作:
//lease:来实现锁自动过期,防止节点down掉后锁不释放造成死锁


/*
1.上锁

2.业务

3.释放锁
3.1.取消自动租约
3.2.释放租约

*/

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse
		keepResp *clientv3.LeaseKeepAliveResponse
		ctx context.Context
		cancelFunc context.CancelFunc
		kv clientv3.KV
		txn clientv3.Txn
		txnResp *clientv3.TxnResponse
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 创建租约(创建租约, 自动续租, 拿着租约去抢占一个key)
	lease = clientv3.NewLease(client)

	// 申请一个5秒的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约的ID
	leaseId = leaseGrantResp.ID

	// 准备一个用于取消自动续租的context
	ctx, cancelFunc = context.WithCancel(context.TODO())

	// defer 会把租约释放掉, 关联的KV就被删除了
	// 确保函数退出后, 自动续租会停止
	defer cancelFunc()
	// 释放租约,立刻释放租约从而立刻删除key
	defer lease.Revoke(context.TODO(), leaseId)

	// 这要这行执行每秒都会向etcd发送保活
	// 5秒后会取消自动续租
	keepRespChan, err = lease.KeepAlive(ctx, leaseId)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约应答的协程
	go func() {
		for {
			select {
				case keepResp = <- keepRespChan:
					if keepRespChan == nil {
						fmt.Println("租约已经失效了")
						goto END
					} else {	// 每秒会续租一次, 所以就会受到一次应答
						fmt.Println("收到自动续租应答:", keepResp.ID)
					}
			}
		}
	END:
	}()

	//  if 不存在key{
	// 		设置它
	// }  else
	// {
	// 		抢锁失败
	// }
	kv = clientv3.NewKV(client)

	// 创建事务
	txn = kv.Txn(context.TODO())

	// 定义事务
	// 如果key不存在
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "xxx", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/lock/job9")) // 否则抢锁失败

	// 提交事务
	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return // 没有问题defer执行
	}

	// 判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用:", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 2, 抢锁成功处理业务
	fmt.Println("处理任务")
	time.Sleep(5 * time.Second)

	// 3, 释放锁(取消自动续租, 释放租约)
	//取消自动续租:假如之前的租约是10 那么这个key在取消自动续约后10秒后过期
	//释放租约:就是强制删除租约 那么这个key立刻删除

}
