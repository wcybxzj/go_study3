package ch32_diy_obj_pool

import (
	"fmt"
	"sync"
	"testing"
)

type Obj1 struct {
	age  int
	name string
	obj2 Obj2
	obj3 *Obj3
}

type Obj2 struct {
	age2  int
	name2 string
}

type Obj3 struct {
	age3  int
	name3 string
}

func f1(ch chan Obj1, wg *sync.WaitGroup) {
	obj1 := <-ch
	obj1.age=101
	obj1.name="ybx1"
	obj1.obj2.age2=201
	obj1.obj2.name2="wc2"
	obj1.obj3.age3=301
	obj1.obj3.name3="ly3"
	ch <- obj1
	wg.Done()
}

func f2(ch chan *Obj1, wg *sync.WaitGroup) {
	obj1 := <-ch
	obj1.age=101
	obj1.name="ybx1"
	obj1.obj2.age2=201
	obj1.obj2.name2="wc2"
	obj1.obj3.age3=301
	obj1.obj3.name3="ly3"
	ch <- obj1
	wg.Done()
}



//测试object在channel中传递情况
/*
输出:
{100 ybx {200 wc} 0xc00000c040}
&{300 ly}
{101 ybx1 {201 wc2} 0xc00000c040}
&{301 ly3}
*/
func TestObjInChannel(t *testing.T) {
	//var channelObj chan Obj1
	channelObj := make(chan Obj1, 1)

	obj1 := Obj1{100, "ybx",
		Obj2{200, "wc"},
		&Obj3{300, "ly"}}

	channelObj  <- obj1

	//
	fmt.Println(obj1)
	fmt.Println(obj1.obj3)

	var wg sync.WaitGroup
	wg.Add(1)
	go f1(channelObj, &wg)
	wg.Wait()
	obj1_1 := <- channelObj

	//
	fmt.Println(obj1_1)
	fmt.Println(obj1.obj3)
}


//测试objectPtr在channel中传递情况
/*
输出:
{100 ybx {200 wc} 0xc00000c040}
&{300 ly}
&{101 ybx1 {201 wc2} 0xc00000c040}
&{301 ly3}
*/
func TestObjPtrInChannel(t *testing.T) {
	//var channelObj chan *Obj1
	channelObj := make(chan *Obj1, 1)

	obj1 := Obj1{100, "ybx",
		Obj2{200, "wc"},
		&Obj3{300, "ly"}}

	channelObj  <- &obj1

	//
	fmt.Println(obj1)
	fmt.Println(obj1.obj3)

	var wg sync.WaitGroup
	wg.Add(1)
	go f2(channelObj, &wg)
	wg.Wait()
	obj1_1 := <- channelObj

	//
	fmt.Println(obj1_1)
	fmt.Println(obj1.obj3)
}

//结论无论在channel里传递的是obj还是objptr都可以修改对象内容
