package __arr_slice

import (
	"fmt"
	"testing"
)

//初始化测试
func TestArrayInit(t *testing.T)  {
	//method1:
	var arr[3]int //声明并且初始化都成0
	for _, v := range arr {
		fmt.Println(v)
	}

	//method2:
	arr1 := [4]int{1, 2, 3, 4}
	
	//method3:
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr1, arr3)
}

func TestArraySection(t *testing.T)  {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:]
	fmt.Printf("%p %p", &arr3, &arr3_sec)
}

func test(arr[3]int)  {
	arr[0] = 111
}

func test1(arr *[3]int) {
	arr[0]=222
}

//数组作为函数参数
func TestArrAsFuncArgs(t *testing.T)  {
	arr := [3]int{11,22,33}
	test(arr)
	t.Log(arr) //[11 22 33]

	test1(&arr)
	t.Log(arr) //[222 22 33]
}

