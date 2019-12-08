package ch17_share_mem

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//fail
func TestCounterUnThreadSafe(t *testing.T)  {
	counter := 0
	for i:=0;i<5000;i++  {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter) //4687
}

//ok
func TestCounterThreadSafe(t *testing.T)  {
	var mut sync.Mutex
	counter := 0
	for i:=0;i<5000;i++  {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter) //5000 耗时1秒
}

//ok
func TestCounterWaitGroup(t *testing.T)  {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i:=0; i<5000; i++  {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println(counter) //5000 耗时0秒
}

