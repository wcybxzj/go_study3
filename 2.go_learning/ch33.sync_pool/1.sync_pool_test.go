package ch33_sync_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//测试1:
//普通对象缓存测试
//sync.pool:可以报错多个对象
//pool.Get:从pool获取对象,如果没有执行New并且返回数据
//pool.Put:是把对象存入到pool
/*
输出:
Create a new object
100
3
*/
func TestSyncPool1(t *testing.T)  {
	pool := &sync.Pool {
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	//1.type assert
	//v := pool.Get().(int)
	//fmt.Println(v)

	//2.type assert
	v, ok := pool.Get().(int)
	if !ok {
		t.Error("type asset fail!")
	}
	fmt.Println(v)

	pool.Put(3)
	//runtime.GC()
	v1, ok := pool.Get().(int)
	if !ok {
		t.Error("type asset fail!")
	}
	fmt.Println(v1)
}

//测试2:
//普通对象缓存测试,在GC后的情况
/*
输出:
Create a new object
100
Create a new object
100
*/
func TestSyncPool2(t *testing.T)  {
	pool := &sync.Pool {
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	v, ok := pool.Get().(int)
	if !ok {
		t.Error("type asset fail!")
	}
	fmt.Println(v)

	pool.Put(3)
	runtime.GC()
	v1, ok := pool.Get().(int)
	if !ok {
		t.Error("type asset fail!")
	}
	fmt.Println(v1)
}

//测试3:
/*
以为pool里根本没有对象,所以每次都要创建
输出:
Create a new object
100
Create a new object
100
Create a new object
100
*/
func TestSyncPool3(t *testing.T)  {
	pool := &sync.Pool {
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	v, _ := pool.Get().(int)
	fmt.Println(v)

	v, _ = pool.Get().(int)
	fmt.Println(v)

	v, _ = pool.Get().(int)
	fmt.Println(v)
}

/*
//输出:
num: 4 val: 100
Create a new object
num: 3 val: 10
num: 2 val: 300
num: 1 val: 200
Create a new object
num: 0 val: 10
*/
func TestSyncPoolInMultiGoroutine(t *testing.T)  {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object ")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(200)
	pool.Put(300)

	var wg sync.WaitGroup
	for i:=0; i<5; i++  {
		wg.Add(1)
		go func(num int) {
			fmt.Println("num:", num, "val:",pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()

}