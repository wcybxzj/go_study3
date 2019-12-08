package main

import (
	"10.go_learing/ch41.micro_kernel/agent"
	"context"
	"errors"
	"fmt"
	"time"
)

type DemoCollector struct {
	evtReceiver agent.EventReceiver
	agtCtx context.Context
	stopChan chan struct{}
	name string
	content string
}

func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector {
		stopChan:    make(chan struct{}),
		name:        name,
		content:     content,
	}
}

func (c *DemoCollector)Init(rec agent.EventReceiver) error {
	fmt.Println("initialize collectos", c.name)
	c.evtReceiver = rec
	return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collector", c.name)
	for {
		select {
		case <-agtCtx.Done():
			c.stopChan <- struct{}{}

		default:
			time.Sleep(time.Millisecond * 50)
			t := time.Now().Format("2006-01-02 15:04:05")
			c.evtReceiver.OnEvent(agent.Event{c.name, c.name+t})
		}
	}
}

func (c *DemoCollector)Stop() error {
	fmt.Println("stop collector", c.name)

	select {
	case <-c.stopChan:
		return nil
	case <-time.After(time.Second):
		return errors.New("failed to stop for timeout")
	}
}

func (c *DemoCollector) Destroy() error {
	fmt.Println(c.name, "release resource.")
	return nil
}

//go run -race agent_test1.go
func main() {
	var err error

	agt := agent.NewAgent(100)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")
	agt.RegisterController("c1", c1)
	agt.RegisterController("c2", c2)

	//应该不会报错
	if err = agt.Start(); err != nil {
		fmt.Println("start1() err:",err.Error())
	}

	if err = agt.Start(); err != nil {
		fmt.Println("start2() err:",err.Error())
	}

	//因为前面已经启动过了会报错
	time.Sleep(time.Second *3)
	agt.Stop()
	agt.Destroy()
}