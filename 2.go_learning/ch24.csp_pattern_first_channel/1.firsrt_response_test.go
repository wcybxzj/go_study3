package ch24_csp_pattern_first_channel

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(time.Millisecond*10)//0.01秒
	return fmt.Sprintf("The result is from %d", i)
}

func FirstResponse(chanNum int) string {
	numOfRunner := 10
	var ch chan string
	if chanNum > 0 {
		ch = make(chan string, chanNum)// buffered channel
	} else {
		ch = make(chan string)// unbuffered channel
	}

	for i:=0; i < numOfRunner ;i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	//return ""
	return <-ch
}

//csp并发模式1:仅需任意任务完成
//1次访问多个http资源, 如果其中一个可以返回立刻获取这个响应

/*
测试1:
输出:
Before: 2(不清楚这两个具体是谁)
The result is from 9(可能的结果是0->9)
After: 11

分析:
本来10个worker协程
其中1个协程写完channel,主协程读取channel后，这个子协程退出了
剩余9个协程, 阻塞在写channel,因为没人在读channel
*/
func TestFirstResponseV1(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse(0))
	/*
	在这个时间窗口里有9个协程因为没人在读取channel而被阻塞不能退出
	这里就等于泄漏了9个协程
	*/

	//这行代码的用处:让那个没被阻塞的的子协程退出
	time.Sleep(time.Second)
	t.Log("After:", runtime.NumGoroutine())
}

/*
测试2:
输出:
Before: 2
The result is from 4
After: 2(没有协程泄漏)
*/
func TestFirstResponseV2(t *testing.T)  {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse(9))
	/*
	在这个时间窗口里有9个协程虽然没人来读取，但是因为channel有9个buffer
	9个子协程没有写没有被阻塞,正常退出了
	*/
	//这行代码的用处:让那个没被阻塞的的子协程退出
	time.Sleep(time.Second)
	t.Log("After:", runtime.NumGoroutine())
}