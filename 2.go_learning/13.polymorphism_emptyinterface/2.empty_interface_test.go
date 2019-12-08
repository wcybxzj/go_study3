package _3_polymorphism_emptyinterface

import (
	"fmt"
	"testing"
)

//mehtod1:
func typeAssert(p interface{})  {
	if v, ok := p.(int); ok{
		fmt.Printf("int:%d\n", v)
		return
	}

	if v, ok := p.(string); ok{
		fmt.Printf("string:%s\n", v)
		return
	}

	fmt.Printf("unknow type\n")
}

//mehtod2:
func switchCase(p interface{})  {
	switch v := p.(type) {
	case int:
		fmt.Printf("int:%d\n", v)
	case string:
		fmt.Printf("string:%s\n", v)
	default:
		fmt.Printf("unknow type\n")
	}
}

func Test(t *testing.T)  {
	typeAssert(10)
	typeAssert("123")
	typeAssert(1.2)

	switchCase(10)
	switchCase("123")
	switchCase(1.2)
}