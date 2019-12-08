package __type_test

import "testing"

type MyInt int64

//golang不支持隐式类型转换
func TestImplict(t *testing.T)  {
	var a int32 =1
	var b int64
	//b = a //error
	b = int64(a)
	var c MyInt
	c = MyInt(a)
	t.Log(a, b, c)
}

//golang不支持指针运算
func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1 //error
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) //int *int
}

//golang 声明的string是空字符串不是nil.而是空字符串
func TestString(t *testing.T) {
	var s string
	t.Log("*"+s+"*")
	t.Log(len(s))
}