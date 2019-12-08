package ch24_csp_pattern_first_channel

import (
	"fmt"
	"testing"
)
//说明:
//在使用unbuffer channel的时，生产者和消费者要同时存在
//假如只有生产者而没有消费者,则生产者会一直阻塞在写channel
//输出:Startfatal error: all goroutines are asleep - deadlock!
func TestOnlyProducer(t *testing.T)  {
	ch := make(chan int)
	fmt.Printf("Start")
	ch<-1
	fmt.Printf("End")
}

//输出:Startfatal error: all goroutines are asleep - deadlock!
func TestOnlyCustomer(t *testing.T)  {
	ch := make(chan int)
	fmt.Printf("Start")
	<-ch
	fmt.Printf("End")
}

