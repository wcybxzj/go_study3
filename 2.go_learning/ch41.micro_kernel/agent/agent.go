package agent

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

const(
	Waitting = iota
	Running
)

var WrongStateError = errors.New("can not operatin in the current state")

//1.在agent中收集collector产生的错误
type CollectorsError struct {
	CollectorsErrors[]error
}

func (ce CollectorsError) Error() string {
	var re []string
	for _, valErr := range ce.CollectorsErrors {
		re = append(re, valErr.Error())
	}
	return strings.Join(re, ";")
}

//2.定义从Collector发向agent的channel中的数据格式
type Event struct {
	Source string
	Content string
}

//3.定义agt中需要实现的借口
type EventReceiver interface {
	OnEvent(evt Event)
}

//4.
type Collector interface {
	Init(recv EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors map[string]Collector
	evtBuf chan Event
	cancel context.CancelFunc
	ctx context.Context
	state int
}

func NewAgent(sizeEvtBuf int) *Agent {
	return &Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan Event, sizeEvtBuf),
		state:      Waitting,
	}
}

func printProcess(events [10]Event)  {
	fmt.Println(events)
}

//TODO:需要优化存在bug
//问题1:如果10改成100就会有问题,因为在main(）中只有3秒时间，collectors写不了怎么多个Event
//办法:借鉴imooc_crontab中worker的LogSink的writeLoop() 没时间改
func (agt *Agent)EventProcess() {
	var events [10]Event
	//defer fmt.Println(events)
	defer printProcess(events)

	for {
		for i:=0; i<10; i++ {
			select {
			case events[i] = <-agt.evtBuf:
			case <-agt.ctx.Done():
				return
			}
		}
		fmt.Println(events)
	}
}

func (agt *Agent)RegisterController(name string, collector Collector) error{
	if agt.state != Waitting {
		return WrongStateError
	}

	agt.collectors[name] = collector
	return collector.Init(agt)
}

func (agt *Agent)startCollectors() error {
	var err error
	var errs CollectorsError
	var mutex sync.Mutex

	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			defer func() {
				mutex.Unlock()
			}()
			mutex.Lock()
			err = collector.Start(ctx)
			if err != nil {
				errs.CollectorsErrors = append(errs.CollectorsErrors,
												errors.New(name+":"+err.Error()))
			}
		}(name, collector, agt.ctx)
	}

	if len(errs.CollectorsErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) stopCollectors() error {
	var err error
	var errs CollectorsError
	for name, colletor := range agt.collectors {
		if err = colletor.Stop(); err != nil {
			errs.CollectorsErrors = append(errs.CollectorsErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorsErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) destroyColletors() error{
	var err error
	var errs CollectorsError
	for name, colletor := range agt.collectors {
		if err = colletor.Destroy(); err != nil {
			errs.CollectorsErrors = append(errs.CollectorsErrors,
				errors.New(name+":"+err.Error()))
		}
	}
	if len(errs.CollectorsErrors) == 0 {
		return nil
	}
	return errs
}

func (agt *Agent) Start() error {
	if agt.state != Waitting {
		return WrongStateError
	}
	agt.state= Running
	agt.ctx, agt.cancel = context.WithCancel(context.Background())
	go agt.EventProcess()
	return agt.startCollectors()
}

func (agt *Agent) Stop() error {
	if agt.state != Running {
		return WrongStateError
	}
	agt.state = Waitting
	agt.cancel()
	return agt.stopCollectors()
}

func (agt *Agent)Destroy() error {
	if agt.state != Waitting {
		return WrongStateError
	}
	return agt.destroyColletors()
}

func (agt *Agent) OnEvent(evt Event) {
	agt.evtBuf <- evt
}