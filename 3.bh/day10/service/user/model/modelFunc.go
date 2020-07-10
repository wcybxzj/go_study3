package model

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"errors"
	"crypto/md5"
	"encoding/hex"
)

// 创建一个全局变量, 代表连接池操作句柄
var RedisPool redis.Pool

// redis连接池初始化
func InitRedis() {
	// redis 连接池
	RedisPool = redis.Pool{
		MaxIdle:         20,
		MaxActive:       50, // 最大存活数 一般 > MaxIdle
		IdleTimeout:     60, // 空闲 超时时间
		MaxConnLifetime: 60 * 5,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.6.108:6379")
		},
	}
}

// 检查图片验证码是否正确 --- imgCode是用户写入的
func CheckImgCode(imgCode, uuid string) bool {
	// 链接 redis
	conn := RedisPool.Get()
	defer conn.Close()

	// 从 redis中 按 uuid 指定 get请求,得到之前存储的 验证码
	code, err := redis.String(conn.Do("get", uuid)) // 直接将 Do方法, 传参给 回复助手函数
	if err != nil {
		fmt.Println("根据uuid查询 redis err:", err)
		return false
	}

	// 将用户输入的验证码 和 get 请求到的验证码 比对. 返回比对结果.
	return imgCode == code
}

// 存储 短信验证码
func SaveSmsCode(phone, code string) error {
	// 从连接池获取一个链接
	conn := RedisPool.Get() // 对应原来的 conn, err := redis.Dial("tcp", "192.168.6.108:6379")
	defer conn.Close()      // 放回

	// 存储 短信验证码,到 redis 中
	_, err := conn.Do("setex", phone+"_code", 60*5, code)

	return err
}

// 校验 短信验证码
func CheckSmsCode(phone, code string) error {
	// 链接 redis 数据库
	conn := RedisPool.Get()

	// 根据 key 获取 value
	smsCode, err := redis.String(conn.Do("get", phone+"_code"))
	if err != nil {
		fmt.Println("根据phone_code获取redis短信验证码失败:", err)
		return err
	}
	// 匹配
	if code != smsCode {
		return errors.New("验证码匹配失败")
	} else {
		fmt.Println("-------------1111-- 验证码匹配成功-----")
	}

	return nil
}

// 注册用户
func RegisterUser(mobile, pwd string) error {
	var user User

	user.Name = mobile
	user.Mobile = mobile

	fmt.Println("-----11111----注册用户------mobile:", mobile)

	m5 := md5.New()                             // 初始化 md5 对象
	m5.Write([]byte(pwd))                       // 将pwd 写入到 缓冲区
	pwd_hash := hex.EncodeToString(m5.Sum(nil)) // 不使用额外的秘钥

	fmt.Println("--------2222----注册用户------pwd:", pwd)

	user.Password_hash = pwd_hash

	fmt.Println("-------------3333---注册用户------:")

	GlobalConn.SingularTable(true)

	fmt.Println("-----------------4444----注册用户------:")

	err := GlobalConn.Create(&user).Error

	fmt.Println("-------------------5555----注册用户------:", err)

	// 使用全局句柄插入到数据库中.
	return err
}