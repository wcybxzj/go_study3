package model

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"encoding/hex"
	"crypto/md5"
	strconv "strconv"
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
	/*	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
		if err != nil {
			fmt.Println("redis.Dial err:", err)
			return false
		}*/
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
	_, err := conn.Do("setex", phone+"_code", 60*3, code)

	return err
}

// 处理用户登录业务, 根据手机号/密码匹配, 返回用户名
func Login(mobile, pwd string) (string, error) {
	// grom 查询 user 表
	var user User

	// 对参数的pwd 求hash
	m5 := md5.New()                             // 初始化 md5 对象
	m5.Write([]byte(pwd))                       // 将pwd 写入到 缓冲区
	pwd_hash := hex.EncodeToString(m5.Sum(nil)) // 不使用额外的秘钥

	// select name from user where mobie=参mobile and password_hash=参pwd哈希
	err := GlobalConn.Select("name").Where("mobile=?", mobile).
		Where("password_hash=?", pwd_hash).Find(&user).Error

	return user.Name, err
}

// 根据用户名获取用户信息
func GetUserInfo(userName string) (User, error) {

	var user User

	err := GlobalConn.Where("name = ?", userName).First(&user).Error

	return user, err
}

// 更新用户名
func UpdateUserName(newName, oldName string) error {

	// update user set name = 'itast40' where name = 旧用户名
	return GlobalConn.Model(new(User)).Where("name = ?", oldName).Update("name", newName).Error
}

// 根据用户名, 更新用户头像
func UpdateAvatar(userName, avatar string) error {
	// SQL : update user set avatar_url = avatar where name = userName
	return GlobalConn.Model(new(User)).Where("name = ?", userName).
		Update("avatar_url", avatar).Error
}

// 保存用户实名信息
func SaveRealName(userName, realName, idCard string) error {
	return GlobalConn.Model(new(User)).Where("name = ?", userName).
		Updates(map[string]interface{}{"real_name": realName, "id_card": idCard}).Error
}


////////////////////////////////////////////////////////////
type HouseInfo struct {
	Acreage   string   `json:"acreage"`
	Address   string   `json:"address"`
	AreaId    string   `json:"area_id"`
	Beds      string   `json:"beds"`
	Capacity  string   `json:"capacity"`
	Deposit   string   `json:"deposit"`
	Facility  []string `json:"facility"`
	MaxDays   string   `json:"max_days"`
	MinDays   string   `json:"min_days"`
	Price     string   `json:"price"`
	RoomCount string   `json:"room_count"`
	Title     string   `json:"title"`
	Unit      string   `json:"unit"`
}
//插入房屋信息到数据库
func AddHouse(userName string, houseInfo HouseInfo) (int, error) {
	//不使用复数表名
	GlobalConn.SingularTable(true)

	//查询用户信息, 主要给下面 UserId 赋值
	var user User

	if err := GlobalConn.Where("name = ?", userName).First(&user).Error; err != nil {
		fmt.Println("获取当前用户信息失败")
		return 0, err
	}
	// 将前端传递来的数据, 写入 House 类对象.
	var house House	// model.go 中的 House类 对应 MySQL 中 house表.

	house.UserId = user.ID
	house.AreaId, _ = strconv.Atoi(houseInfo.AreaId)

	house.Address = houseInfo.Address
	house.Title = houseInfo.Title

	house.Room_count, _ = strconv.Atoi(houseInfo.RoomCount)
	house.Acreage, _ = strconv.Atoi(houseInfo.Acreage)
	house.Price, _ = strconv.Atoi(houseInfo.Price)
	house.Unit = houseInfo.Unit
	house.Capacity, _ = strconv.Atoi(houseInfo.Capacity)
	house.Beds = houseInfo.Beds
	house.Deposit, _ = strconv.Atoi(houseInfo.Deposit)
	house.Min_days, _ = strconv.Atoi(houseInfo.MinDays)
	house.Max_days, _ = strconv.Atoi(houseInfo.MaxDays)

	// 遍历家具设施 id.
	for _, fid := range houseInfo.Facility {
		id, _ := strconv.Atoi(fid)

		//获取id对应的家具对象
		var fac Facility
		fac.Id = id

		GlobalConn.First(&fac)
		house.Facilities = append(house.Facilities, &fac)
	}

	// 将 前端传来的数据, 插入到 MySQL - house 表中
	if err := GlobalConn.Create(&house).Error; err != nil {
		fmt.Println("插入房屋信息失败")
		return 0, err
	}
	// 如果插入成功, house 表 gorm.Model 字段会产生 ID值.

	fmt.Println(house)  // 打印看看

	// <接口文档.doc> 插入成功,返回 house_id
	return int(house.ID), nil
}