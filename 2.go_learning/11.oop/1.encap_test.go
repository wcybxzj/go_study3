package _1_oop

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployeeObj(t *testing.T)  {
	//method1
	e := Employee{"0", "Bob", 20}

	//method2
	e1 := Employee{Name:"Mike"}

	//method3:返回的是指针类型
	e2 := new(Employee)//注意这⾥里里返回的引⽤用/指针，相当于 e := &Employee{}
	e2.Id="111"
	e2.Age=22
	e2.Name="Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e2)

	t.Logf("e is %T", e)
	t.Logf("e is %T", e2)

}

//不推荐
//使用e Employee 存在结构体复制
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) //c0000c0010
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//推荐!!!!!!!!!!:
//使用e *Employee不存在结构体复制
func (e *Employee) String1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) //c000072370
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T)  {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Adderss is %x\n", unsafe.Pointer(&e.Name)) //c000072370
	t.Log(e.String())
	t.Log(e.String1())
}