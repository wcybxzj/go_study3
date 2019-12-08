package ch20_channel_close

import (
	"fmt"
	"testing"
	"time"
)

/*
输出:
123
start close channel
close close channel
channel close, can not reads
*/
func TestCloseNoneBufferChannel(t *testing.T)  {
	ch := make(chan int)
	go func() {
		ch<-123
		fmt.Println("start close channel")
		close(ch)
		fmt.Println("close close channel")
	}()

	go func() {
		time.Sleep(time.Second)
		//fmt.Println("func2:")
		for {
			if v, ok:=<-ch; ok {
				fmt.Println(v)
			}else {
				fmt.Println("channel close, can not reads")
				return
			}
		}
	}()

	time.Sleep(time.Second*3)
}

/*
输出:
start close channel
close close channel
123
channel close, can not reads
可以看到即使channel关闭了,在消费者还是能读取到数据
*/
func TestCloseBufferChannel(t *testing.T)  {
	ch := make(chan int, 1)
	go func() {
		ch<-123
		fmt.Println("start close channel")
		close(ch)
		fmt.Println("close close channel")
	}()

	go func() {
		time.Sleep(time.Second)
		//fmt.Println("func2:")
		for  {
			if v, ok:=<-ch; ok {
				fmt.Println(v)
			}else {
				fmt.Println("channel close, can not reads")
				return
			}
		}
	}()

	time.Sleep(time.Second*3)
}
