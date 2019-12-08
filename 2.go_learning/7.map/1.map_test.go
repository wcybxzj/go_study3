package main

import "testing"

func Test(t *testing.T) {
	//method1:
	m1 := map[int]int{1: 1, 2: 4, 3:9}
	t.Logf("len m1=%d", len(m1)) //3

	//method2:
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))//1

	//method3:
	m3 :=make(map[int]int, 10) //10是capacity
	t.Logf("len m3=%d", len(m3))//0
}

//go语言对key不存在的map的value会保存一个默认值
func TestAccessNotExistingKey(t *testing.T)  {
	m1 := map[int]int{}
	t.Log(m1[1]) //访问一个不存在的key返回了0

	m1[2] = 0
	t.Log(m1[2]) //访问一个赋值为0 的key也返回了0

	//解决办法:
	if v, ok :=m1[1]; ok {
		t.Logf("m1[1] is exists value is %d", v)
	} else {
		t.Logf("m1[1] is not exists ")
	}

	if v, ok :=m1[2]; ok {
		t.Logf("m1[2] is exists value is %d", v)
	} else {
		t.Logf("m1[2] is not exists ")
	}
}