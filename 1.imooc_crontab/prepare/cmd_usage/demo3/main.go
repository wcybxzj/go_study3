package main

import (
	"os/exec"
	"context"
	"time"
	"fmt"
)

type result struct {
	err error
	output []byte
}
//例子3:
//执行1个cmd,让它在一个协程里去执行, 让它执行2秒: sleep 2; echo hello;
//1秒的时候, 我们杀死cmd

//以前我为实现一样的功能写了1个RunCmdAsyncWithTimeout()方法实现golang执行bash带超时功能
//其实根本不用怎么麻烦 用exec.CommandContext()就能实现一样的功能
func main() {

	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)

	//创建了一个结果队列用于读取子协程的输出
	resultChan = make(chan *result, 1000)

	// 创建上下文
	// context:   chan byte
	// cancelFunc:  close(chan byte)
	ctx, cancelFunc = context.WithCancel(context.TODO())

	//创建执行命令的协程
	go func() {
		var (
			output []byte
			err error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello;")

		// 执行任务, 捕获输出
		output, err = cmd.CombinedOutput()

		// 把任务输出结果, 传给main协程
		resultChan <- &result{
			err: err,
			output: output,
		}
	}()

	// 继续往下走
	time.Sleep(1 * time.Second)

	// 取消上下文
	cancelFunc()

	// 在main协程里, 等待子协程的退出，并打印任务执行结果
	res = <- resultChan

	// 打印任务执行结果
	fmt.Println(res.err, string(res.output))
}
