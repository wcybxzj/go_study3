package main_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

//test1: 不合理的类型转换
/*
输出:
0xc00001a120
5e-323 //虽然强转了但是数不对
*/
func Test1(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

type MyInt int

//test2:合理的类型转换
func Test2(t *testing.T) {
	a:= []int{1, 2 ,3, 4}
	b:= *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

//test3:atomic测试
func Test3(t *testing.T) {
	var sharePtr unsafe.Pointer
	wFn := func() {
		data := []int{}
		for i:=0; i<100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&sharePtr,unsafe.Pointer(&data))
	}

	rFn := func() {
		data := atomic.LoadPointer(&sharePtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				wFn()
				time.Sleep(time.Microsecond*100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i:=0;i<10 ;i++  {
				rFn()
				time.Sleep(time.Microsecond*100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

//test4:不使用atomic测试
func Test4(t *testing.T) {
	data := []int{}

	//var sharePtr unsafe.Pointer
	wFn := func() {
		for i:=0; i<100; i++ {
			data = append(data, i)
		}
	}

	rFn := func() {
		fmt.Println(&data, data)
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				wFn()
				time.Sleep(time.Microsecond*100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i:=0;i<10 ;i++  {
				rFn()
				time.Sleep(time.Microsecond*100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}