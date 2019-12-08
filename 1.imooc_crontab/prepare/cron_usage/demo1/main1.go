package main

import (
	"github.com/gorhill/cronexpr"
	"fmt"
	"time"
)

//用linux crontab格式简单测试下golang的cronexpr库获取定时的目标时间
func main() {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)

	// linux crontab
	// 秒粒度, 年配置(2018-2099)
	// 哪一分钟（0-59），哪小时（0-23），哪天（1-31），哪月（1-12），星期几（0-6）

	//重点: */5 * * * * * *
	//的意思是每小时的5,10,15,20,25,30,35,40,45,50,55,00
	//而不是设置后的每5分钟执行
	//例如你设置时候是6点3分 下次执行的时间就是6点5分


	//cronexpr.Parse():判断输入的计划任务是否正确 正确返回expr对象,错误返回err
	//cronexpr.MustParse():认为不用判断输入的计划任务是否正确直接返回expr对象
	if expr, err = cronexpr.Parse("*/5 * * * * "); err != nil {
		fmt.Println(err)
		return
	}

	// 0, 6, 12, 18, .. 48..

	// 当前时间
	now = time.Now()

	// 下次调度时间
	nextTime = expr.Next(now)

	fmt.Println(now, nextTime)
}
