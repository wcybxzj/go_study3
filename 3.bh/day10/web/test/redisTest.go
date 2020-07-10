package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {

	// 链接数据库
	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 操作数据库
	reply, err := conn.Do("set", "itcast", "itheima")
	if err != nil {
		fmt.Println("Do err:", err)
		return
	}

	// 使用 回复助手函数
	str, e := redis.String(reply, err)

	fmt.Println(str, e)
}
