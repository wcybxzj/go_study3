package ch32_diy_obj_pool

import (
	"fmt"
	"testing"
	"time"
)

type Student struct {
	age int
	name string
}

//归还对象测试
func Test1(t *testing.T)  {
	pool := NewPool(10)

	//对一个已经满了的pool放入对象
	if err := pool.ReleaseObj(&ReuseableObj{}); err != nil {
			t.Error(err)
		}
}

//申请对象测试
func Test2(t *testing.T) {
	pool := NewPool(10)
	for i:=0; i<11; i++ {
		if v, err := pool.GetObj(time.Second); err != nil {
			t.Error(err)
		}else{
			fmt.Printf("%T\n", v)
			/*
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
			*/
		}
	}
}