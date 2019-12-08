package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

const WATCH_DIR  = "/cron/jobs/job7"

//watch
//使用GO API比较麻烦还需要获取revision+key才能进行watch
//1个协程写入删除一个key
//主协程watch这个key
//主协程给watcher的channel设置了5秒超时
func main() {
	var (
		config             clientv3.Config
		client             *clientv3.Client
		err                error
		kv                 clientv3.KV
		watcher            clientv3.Watcher
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		watchRespChan      <-chan clientv3.WatchResponse
		watchResp          clientv3.WatchResponse
		event              *clientv3.Event
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

	// KV
	kv = clientv3.NewKV(client)

	// 模拟etcd中KV的变化
	// 生成1个key 然后删除它
	go func() {
		for {
			val := time.Now().Format("2006-01-02 15:04:05")
			kv.Put(context.TODO(), WATCH_DIR, val)
			kv.Delete(context.TODO(), WATCH_DIR)
			time.Sleep(1 * time.Second)
		}
	}()

	// 先GET到当前的值，并监听后续变化
	getResp, err = kv.Get(context.TODO(), WATCH_DIR)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 现在key是存在的
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值:", string(getResp.Kvs[0].Value))
	}

	// 从下一个事务ID进行监控
	// 当前etcd集群事务ID, 单调递增的
	watchStartRevision = getResp.Header.Revision + 1

	// 创建一个watcher
	watcher = clientv3.NewWatcher(client)

	// 启动监听
	fmt.Println("从该版本向后监听:", watchStartRevision)

	ctx, cancelFunc := context.WithCancel(context.TODO())

	//只能执行一次, 5秒钟后执行
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})

	//监控 WATCH_DIR 和 WATCH_DIR下面的子目录
	//watchRespChan = watcher.Watch(ctx, WATCH_DIR, clientv3.WithRev(watchStartRevision), clientv3.WithFromKey())

	//只能监控 WATCH_DIR
	watchRespChan = watcher.Watch(ctx, WATCH_DIR, clientv3.WithRev(watchStartRevision))

	// 处理kv变化事件
	//CreateRevision:key创建的revision
	//ModRevision:key最后修改的revision
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				fmt.Println("修改为:", string(event.Kv.Value),
					"CreateRevision(key第一次创建这个key时的版本):", event.Kv.CreateRevision,
					"ModRevision(key最后一次修改的版本):", event.Kv.ModRevision)
			case clientv3.EventTypeDelete:
				fmt.Println("删除了", "ModRevision:", event.Kv.ModRevision)
			}
		}
	}

	fmt.Println("main goroutine finished!")
}
