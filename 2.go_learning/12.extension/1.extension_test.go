package _2_extension

import (
	"fmt"
	"testing"
)

//父类型
type Pet struct{

}

func (p *Pet) Speak(){
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

//子类型
type Dog struct{
	Pet //匿名内嵌类型
}

//问题:非常不合理好麻烦
//这个函数必须写重写一遍即使里面的内容和父类型中的函数是一样的
func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}


func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
