package _0_func

import (
	"fmt"
	"testing"
)

func Sum(ops ...int) int{
	ret := 0
	for _, v := range ops  {
		//ret = ret +v
		ret += v
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(11,22,33))
}

func Clear()  {
	fmt.Println("Clear resources.111111111111111")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("pppppppppppppppppp")//即使有panic defer然后会执行

}