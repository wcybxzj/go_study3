package __map_ext

import (
	"testing"
)

//map的值是函数,来实现工厂模式
func TestMapWithFunValue(t *testing.T)  {
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func isSet(n int, mySet map[int]bool, t *testing.T)  {
	//mySet[1]=false
	if mySet[n] {
		t.Logf("n%d is exsiting", n)
	} else {
		t.Logf("n%d is not exsiting", n)
	}


}

//用map模拟set类型
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	t.Log("len:",len(mySet)) // 0
	mySet[1] = true
	t.Log("len:",len(mySet)) // 1
	isSet(1, mySet, t)

	mySet[3] = true
	t.Log(len(mySet)) //2
	delete(mySet, 1)
	isSet(1, mySet, t)
}
