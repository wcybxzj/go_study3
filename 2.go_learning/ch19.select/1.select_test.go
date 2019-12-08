package ch19_select

import (
	"fmt"
	"testing"
	"time"
)

func Job1() string {
	time.Sleep(time.Second/2)
	return "Job1 Done"
}

func asyncJob1UseNonBLockChannel() chan string {
	ch := make(chan string,1) //这里channel缓冲1
	go func() {
		ret := Job1()
		fmt.Println("Job1 work finish!")
		ch <- ret
		fmt.Println("Job1 real finish!")
	}()
	return ch
}

//用select为异步任务加一个超时机制
func TestSelectTimeOut(t *testing.T)  {
	select {
		case ret := <-asyncJob1UseNonBLockChannel():
			fmt.Println(ret)
		case <- time.After(time.Microsecond * 100):
			fmt.Println("time out!")
	}
}