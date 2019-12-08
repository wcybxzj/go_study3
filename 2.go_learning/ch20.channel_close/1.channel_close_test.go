package ch20_channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducerV1(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10 ;i++  {
			ch <- i
		}
		wg.Done()
	}()
}

func dataReceiverV1(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10 ;i++  {
			data := <- ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func dataProducerV3(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<20 ;i++  {
			ch <- i
		}
		close(ch)
		//ch <-11 //panic: send on closed channel
		wg.Done()
	}()
}

// v, ok <-ch; ok 为 bool 值，true 表示正常接受，false 表示通道关闭
func dataReceiverV3(ch chan int, wg *sync.WaitGroup)  {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			}else{
				break
			}
		}
		wg.Done()
	}()
}

//v1:
//TestV1()
//消费者为了能知道生产多少才结束,消费者和生产者的消费数量是一样都写死成了10

//v2:ch21通过close channel的广播机制来解决这个问题
//生产者在channel中传递一个token例如-1,来让消费者知道生产结束了
//但是问题是如果有10个消费者,生产者还要关心消费者的数量发送10个-1

//v3:TestV3()
//当生产者消费完之后就close channel,这个信息将会广播到所有的消费者
//这样即使有100个消费者在从chnnel读取数据,也立刻知道这个情况可以结束当前gorutine

func TestV1(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducerV1(ch ,&wg)
	wg.Add(1)
	dataReceiverV1(ch, &wg)
	wg.Wait()
}

func TestV3(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	//1个producer
	wg.Add(1)
	dataProducerV3(ch ,&wg)
	//3个消费者
	wg.Add(1)
	dataReceiverV3(ch, &wg)
	wg.Add(1)
	dataReceiverV3(ch, &wg)
	wg.Add(1)
	dataReceiverV3(ch, &wg)
	wg.Wait()
}