package ch23_once

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

var once sync.Once
var singleInstance * Singleton

type Singleton struct {
	data string
}

func GetSingletonObj() *Singleton {
	once.Do(func() {
			fmt.Println("Create Obj")
			singleInstance = new(Singleton)
		})
	return singleInstance
}

/*
输出:
Create Obj
obj address:C000050180
obj address:C000050180
obj address:C000050180
obj address:C000050180
obj address:C000050180
obj address:C000050180
obj address:C000050180
obj address:C000050180
obj address:C000050180
obj address:C000050180

once其实功能可以用package中的init()来实现
*/
func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("obj address:%X\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}