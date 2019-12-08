package __constant

import "testing"

//1,1,2,3,5,8,13....
func TestFib(t *testing.T) {
	a:=1
	b:=1
	t.Log(a)

	for i:=0; i<5; i++  {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + b
	}
}

//1次交换2个变量
func TestExchange(t *testing.T) {
	a:=1
	b:=2

	a, b = b, a
	t.Log(a, b)
}

