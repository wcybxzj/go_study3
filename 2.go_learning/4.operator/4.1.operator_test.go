package __operator

import (
	"testing"
)

// 1.比较数组
// ⽤用 == ⽐比较数组
// 相同维数且含有相同个数元素的数组才可以⽐比较
// 每个元素都相同的才相等
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

/*
2. &^ 按位置零
与其他主要编程语⾔言的差异
左边  右边 结果
1  &^ 0   1
0  &^ 0   0
1  &^ 1   0
0  &^ 1   0

右边只要是0, 左边的的哪一位保持不变
右边只要是1, 左边的的哪一位被清零
*/

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 //0111
	a = a &^ Readable
	t.Log(a) //6

	a = a &^ Executable
	t.Log(a) //2

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable) //false true false
}
