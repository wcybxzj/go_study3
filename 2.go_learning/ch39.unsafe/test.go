package main

import "fmt"

var a *int

type keysDataSt struct {
	KsInfos int
}

var KsDataGlobal *keysDataSt

//test1:测试int在并发冲突下的情况
func test1() {
	b := 123
	go func() {
		for {
			a = &b
			fmt.Println(*a)
		}
	}()

	go func() {
		for  {
			fmt.Println(*a)
		}
	}()
}

//test2:测试struct在并发冲突下的情况
//直接说明了cto Willan说的用地址复制就不会冲突是错的
func test2() {
	var ksData keysDataSt
	go func() {
		ksData.KsInfos=123 //冲突
		KsDataGlobal = &ksData
	}()

	go func() {
		fmt.Println(ksData.KsInfos)//冲突
	}()
}

//go run -race test.go > /tmp/123
func main() {
	//test1()
	test2()
}
