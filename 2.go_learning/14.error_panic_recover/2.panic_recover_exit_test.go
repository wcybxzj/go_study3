package _4_error_panic_recover

import (
	"fmt"
	"os"
	"testing"
)

//exit:不会执行defer的 不打印堆栈信息
//panic:执行defer 并且打印堆栈信息
func TestExit(t *testing.T) {
	defer func() {
		fmt.Println("不执行")
	}()

	fmt.Println("Start")
	os.Exit(-1)
}

func TestPanic(t *testing.T)  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[[[[recovered from]]]]]", err)
		}
	}()
	fmt.Println("Start")


	panic("1111")

	//panic(errors.New("Something wrong!"))
}