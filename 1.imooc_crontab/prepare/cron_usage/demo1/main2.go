package main

import (
	"github.com/gorhill/cronexpr"
	"fmt"
	"time"
)

//golang 实现5秒后调用1个程序
func main() {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)

	//golang crontab 支持7位设置计划任务时间
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 0, 6, 12, 18, .. 48..

	// 当前时间
	now = time.Now()

	// 下次调度时间
	nextTime = expr.Next(now)

	//fmt.Println(now, nextTime)

	// 等待这个定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了:", nextTime)
	})

	//防止还没到达定时时间 主协程退出
	time.Sleep(5 * time.Second)
}
