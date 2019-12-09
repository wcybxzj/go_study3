package agent

import (
	"context"
	"errors"
	"fmt"
	"time"

	//"github.com/sasha-s/go-deadlock"
	"strings"
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

// 事件批次
type EventBatch struct {
	Events []* Event	// 多条事件
}

//3.定义agt中需要实现的借口
type EventReceiver interface {
	OnEvent(evt *Event)
}

//4.
type Collector interface {
	Init(recv EventReceiver) error
	Start(agtCtx context.Context) error
	Stop() error
	Destroy() error
}

type Agent struct {
	collectors 	  map[string]Collector
	evtBuf 		  chan *Event
	evtBufTimeOut chan *EventBatch
	cancel context.CancelFunc
	ctx context.Context
	state int
	batchMax int
}

func NewAgent(sizeEvtBuf int, batchMax int) *Agent {
	return &Agent{
		collectors: map[string]Collector{},
		evtBuf:     make(chan *Event, sizeEvtBuf),
		evtBufTimeOut:     make(chan *EventBatch, sizeEvtBuf),
		state:      Waitting,
		batchMax:	batchMax,
	}
}

func printProcess(events []*Event) {
	for _, ptr := range events {
		fmt.Print( ptr.Source, "-", ptr.Content, " ")
	}
	fmt.Println()
}

//TODO:需要优化存在bug
//问题1:如果10改成100就会有问题,因为在main(）中只有3秒时间，collectors写不了怎么多个Event
//办法:借鉴imooc_crontab中worker的LogSink的writeLoop()
func (agt *Agent)EventProcessOld() {
	var events [10]*Event
	//defer fmt.Println(events)
	//defer printProcess(events)

	for {
		for i:=0; i<10; i++ {
			select {
			case events[i] = <-agt.evtBuf:
			case <-agt.ctx.Done():
				return
			}
		}

		for _,ptr := range events{
			fmt.Print( ptr.Source, "-", ptr.Content, " ")
		}
		fmt.Println()
	}
}

//优化版:
func (agt *Agent)EventProcess() {
	var (
		event *Event
		eventBatch *EventBatch //正常批次
		timeEventBatch *EventBatch //超时批次
		commitTimer *time.Timer
	)

	for {
		select {
		case event = <- agt.evtBuf:
			if eventBatch == nil {
				eventBatch = &EventBatch{}
				//1秒执行一次
				commitTimer = time.AfterFunc(
					time.Duration(50 * time.Millisecond),
					func(batch *EventBatch, agt *Agent) func() {
						return func() {
							agt.evtBufTimeOut <- batch
						}
					}(eventBatch, agt),
				)
			}
			eventBatch.Events = append(eventBatch.Events, event)

			if len(eventBatch.Events) >= agt.batchMax {
				printProcess(eventBatch.Events)
				eventBatch = nil
				commitTimer.Stop()
			}
		case timeEventBatch = <- agt.evtBufTimeOut:
			if timeEventBatch != eventBatch {
				continue
			}
			printProcess(timeEventBatch.Events)
			eventBatch = nil
		}
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

	//var mutex deadlock.Mutex
	//var mutex sync.Mutex

	for name, collector := range agt.collectors {
		go func(name string, collector Collector, ctx context.Context) {
			/*
			defer func() {
				mutex.Unlock()
			}()
			*/

			//mutex.Lock() //如果放这里只能有一个collector运行因为这个协程不会去自动退出
			err = collector.Start(ctx)
			//mutex.Lock() //根本不执行,多余弄什么mutex

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

func (agt *Agent) OnEvent(evt *Event) {
	agt.evtBuf <- evt
}