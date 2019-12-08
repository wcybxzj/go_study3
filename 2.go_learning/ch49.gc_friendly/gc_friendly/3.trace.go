package main

import (
	"os"
	"runtime/trace"
)

func test()  {
	for i:=0; i<1000; i++ {

	}
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if  err!=nil{
		panic(err)
	}
	defer trace.Stop()

	test()
}
