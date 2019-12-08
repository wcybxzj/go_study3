package ch18_csp

import (
	"fmt"
	"testing"
	"time"
)

func Job1() string {
	time.Sleep(time.Second/2)
	return "Job1 Done"
}

func Job2() {
	fmt.Println("Job2 start")
	time.Sleep(time.Second)
	fmt.Println("Job2 end")
}

//输出:
//Job1 Done
//Job2 start
//Job2 end
func Test1(t *testing.T) {
	fmt.Println(Job1())
	Job2()
}

func asyncJob1UseBLockChannel() chan string {
	ch := make(chan string)
	go func() {
		ret := Job1()
		fmt.Println("Job1 work finish!")
		ch <- ret
		fmt.Println("Job1 real finish!")
	}()
	return ch
}

/*
输出:
Job2 start
Job1 work finish!
Job2 end <-----job2完成后从channel读取数据后job1的goroutie才能结束
Job1 Done
Job1 real finish!
*/
func Test2(t *testing.T) {
	ch := asyncJob1UseBLockChannel()
	Job2()
	fmt.Println(<-ch)
}

func asyncJob1UseNonBLockChannel() chan string {
	ch := make(chan string,1) //这里channel缓冲1
	go func() {
		ret := Job1()
		fmt.Println("Job1 work finish!")
		ch <- ret
		fmt.Println("Job1 real finish!")
	}()
	return ch
}

/*
输出:
Job2 start
Job1 work finish!
Job1 real finish! <--job1不用在等待job2是否从channel读取
Job2 end
Job1 Done
*/
func Test3(t *testing.T) {
	ch := asyncJob1UseNonBLockChannel()
	Job2()
	fmt.Println(<-ch)
}
