package _0_func

import (
	"fmt"
	"testing"
	"time"
)

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

//函数作为参数和返回值，实现一个装饰器模式
func TestFn(t *testing.T)  {
	fn := timeSpent(slowFun)
	ret := fn(100)
	t.Log(ret) //100
}