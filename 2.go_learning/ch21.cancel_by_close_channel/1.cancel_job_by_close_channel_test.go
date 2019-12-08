package ch21_cancel_by_close_channel

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
		case <-cancelChan:
			return true
		default:
			return false
	}
}

func cancel_v1(cancelChan chan struct{})  {
	cancelChan <- struct{}{}
}

func cancel_v2(cancelChan chan struct{})  {
	close(cancelChan)
}

func doTest(t *testing.T, diyFunc func (cancelChan chan struct{})) {
	cancelChan := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Microsecond)
			}
			fmt.Println("goroutine i:", i, "Cancelled")
		}(i, cancelChan)
	}
	diyFunc(cancelChan)
	time.Sleep(time.Second * 1)
}

//输出:启动了5个go 只有1个被正产关闭了
//goroutine i: 4 Cancelled
func TestV1(t *testing.T) {
	doTest(t, cancel_v1)
}

/*
select+close channel的广播机制来实现取消goroutine
输出:
goroutine i: 3 Cancelled
goroutine i: 1 Cancelled
goroutine i: 4 Cancelled
goroutine i: 2 Cancelled
goroutine i: 0 Cancelled
*/
func TestV2(t *testing.T) {
	doTest(t, cancel_v2)
}
