package pb

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 要求, 服务端在注册 rpc 对象时, 能让编译器检测出,注册对象是否合法.
type MyInterface interface {
	HelloWorld(string, *string) error  // 创建接口, 在接口中, 指定方法的原型.
}

// 创建类
type World struct {
}

// 给类绑定方法
func (this *World)HelloWorld(name string, resp *string) error {
	*resp = name + "你好"
	return nil
	//return errors.New("未知的错误!")
}

// 调用该方法, 需要给 i 传,实现了 HelloWorld 方法的 类对象.
func RigsterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

////////////////////// 以上是服务端封装, 以下是 客户端封装 ///////////////////////////////

//像调用本地函数一样,调用远程函数
type MyClient struct {
	c *rpc.Client	 // 根据文档中 Call方法, 给定类型
}

// 由于使用了 c 调用 Call方法, 因此,c需要初始化.
func InitClient(addr string) MyClient {

	conn, _ := jsonrpc.Dial("tcp", addr)

	return MyClient{c : conn}
}

// 此函数原型,也是参照 上面的 Interface 而来.
func (this *MyClient)HelloWorld(a string, b *string) error {

	// 参1,参照 Interface 和 RegisterName 填写.
	return this.c.Call("hello.HelloWorld", a, b)
}
