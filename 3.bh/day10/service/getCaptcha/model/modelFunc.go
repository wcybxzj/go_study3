package model

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

// 存储 图片验证码到 redis
func SaveImgCode(code, uuid string) error {
	// 链接数据库
	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
	if err != nil {
		fmt.Println("Dial err:", err)
		return err
	}
	defer conn.Close()

	// 操作数据库 --- 指定存储 5 分钟.
	_, err = conn.Do("setex", uuid, 60*5, code)

	return err
}
