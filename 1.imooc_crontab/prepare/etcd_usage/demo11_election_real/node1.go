package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

func main() {
	var (
		config clientv3.Config
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	cli, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// create two separate sessions for election competition
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	e1 := concurrency.NewElection(s1, "/my-election/")

	// create competing candidates, with e1 initially losing to e2
	var wg sync.WaitGroup
	wg.Add(1)
	electc := make(chan *concurrency.Election, 2)
	go func() {
		defer wg.Done()
		// delay candidacy so e2 wins first
		time.Sleep(3 * time.Second)
		if err := e1.Campaign(context.Background(), "e1"); err != nil {
			log.Fatal(err)
		}
		electc <- e1
	}()

	cctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	e := <-electc
	fmt.Println("completed first election with", string((<-e.Observe(cctx)).Kvs[0].Value))

	// resign so next candidate can be elected
	//if err := e.Resign(context.TODO()); err != nil {
	//	log.Fatal(err)
	//}

	wg.Wait()

	for i:=1; i<30 ;i++ {
		time.Sleep(time.Second)
		fmt.Println(strconv.Itoa(i))
	}
}
