package _0_func

import (
	"fmt"
	"testing"
)

func func1(arr *[]int)  {
	*arr = append(*arr,123)
}

func TestArr(t *testing.T) {
	var arr[]int
	func1(&arr)
	fmt.Println(arr)
}

func test2(m map[int]int)  {
	m[100]=200
}

func TestMap(t *testing.T) {
	m:=map[int]int{}
	test2(m)
	fmt.Println(m)
}