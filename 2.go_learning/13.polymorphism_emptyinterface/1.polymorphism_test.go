package _3_polymorphism_emptyinterface

import (
	"fmt"
	"testing"
)

type Code string

type Progammer interface {
	WriteHelloWorld() Code
}

type Go struct {

}

func (p *Go) WriteHelloWorld()Code {
	return "golang"
}

type Java struct {

}

func (p *Java) WriteHelloWorld()Code {
	return "java"
}

//接口做为函数参数必须传指针类型
func write(p Progammer)  {
	//%T:golang类型
	//%v:用数据的默认格式
	fmt.Printf("%T %v\n",p ,p.WriteHelloWorld())
}

func TestPloymorphsim(t *testing.T) {
	//goProgrammer := Go{} //error
	goProgrammer := &Go{}
	JavaProgrammer := new(Java)
	write(goProgrammer)
	write(JavaProgrammer)
}