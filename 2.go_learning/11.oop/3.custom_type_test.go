package _1_oop

import (
	"fmt"
	"testing"
	"time"
)

type CustomTYpe  func(op int) int

func timeSpent(inner CustomTYpe) CustomTYpe {
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}