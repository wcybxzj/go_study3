package main

import (
	"github.com/jinzhu/gorm"
 	_ "github.com/go-sql-driver/mysql"   // 前面的 下划线 不能省略
	"fmt"
	"time"
	"crypto/md5"
	"encoding/hex"
	"bj40ihome/web/model"
)

// 创建一个全局结构体, 映射数据库中的表.
type Student struct {
	gorm.Model	// 匿名成员. ----继承
	// Id int		// 被自动指定为主键. 提高查询速度.
	Name string
	Age int
}

type Teacher struct {
	Id int
	Name string `gorm:"size:100;default:'xiaoming'"`
	Age int
	Addr string `gorm:"size:80;default:'北京朝阳'"`  // 新增
	Class int 	`gorm:"not null"`    // 新增
	Join time.Time	`gorm:"type:date"`
	His time.Time
}

// 创建 全局的链接句柄
var GlobalConn *gorm.DB

func main() {
	// 链接数据库 --- 指定驱动
	conn, err := gorm.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/ihome40?parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("链接数据库失败...", err)
		return
	}

	conn.DB().SetMaxIdleConns(10)	// 初始数量
	conn.DB().SetMaxOpenConns(100)  // 最大数

	GlobalConn = conn

	// 指定不使用复数的表 --- 创建/使用表,都要指定
	GlobalConn.SingularTable(true)

	// 借助 gorm 创建数据库表  --- AutoMigrate(
	//fmt.Println(GlobalConn.AutoMigrate(new(Student)).Error)

	fmt.Println(GlobalConn.AutoMigrate(new(Teacher)).Error)

	// 测试插入
	InsertData()

	// 测试查询
	//SearchData()

	// 测试更新
	//UpdateData()

	// 测试软删除
	//DeleteData()

}

// 插入数据
func InsertData() {

	var user model.User

	user.Name = "18610382737"
	user.Mobile = "18610382737"

	m5 := md5.New()                             // 初始化 md5 对象
	m5.Write([]byte("123"))                       // 将pwd 写入到 缓冲区
	pwd_hash := hex.EncodeToString(m5.Sum(nil)) // 不使用额外的秘钥

	user.Password_hash = pwd_hash

	// 使用全局句柄插入到数据库中.
	err := GlobalConn.Create(&user).Error
	fmt.Println(err)
}

// 查询数据
func SearchData()  {
	//var stu Student

	//GlobalConn.First(&stu)  // select * from student order by id limit 1
	//err := GlobalConn.Select("name, age").First(&stu).Error   //获取一部分字段
	//err := GlobalConn.Select("name", "age").First(&stu).Error  // 错误!!
	//err := GlobalConn.Select("name, age").Last(&stu).Error

	// 查询所有数据
/*	var stu []Student

	// select 放到 Find 后,语法无误, 没有效果
	err := GlobalConn.Select("name, age").Find(&stu).Error

	fmt.Println(stu)

	fmt.Println(err)*/

	var stu []Student

	// select name, age from student where name = 'lisi';
	//GlobalConn.Select("name, age").Where("name = ?", "lisi").Find(&stu)

	// select * from student where name = 'lisi' and age = 22;
	// GlobalConn.Where("name = ?", "lisi"). Where("age = ?", 18).Find(&stu)
	// GlobalConn.Where("name = ? and age = ?", "lisi", 18).Find(&stu)
	GlobalConn.Unscoped().Find(&stu)  // 查询软删除后的数据

	fmt.Println(stu)
}

// 更新数据
func UpdateData()  {

/*	stu.Id = 5
	stu.Name = "liu7"
	fmt.Println(GlobalConn.Save(&stu).Error)*/

/*	var stu Student

	//err := GlobalConn.Model(new(Student)).Where("name = ?", "张安").
	//	Update("name", "zhang33").Error
	err := GlobalConn.Model(&stu).Where("name = ?", "AAA").
		Update("name", "zhangAAA").Error*/

	err := GlobalConn.Model(new(Student)).Where("id = ?", 1).
		Updates(map[string]interface{}{"name":"陈8","age":119}).Error

	fmt.Println(err)
}

// 删除数据
func DeleteData()  {
	// delete from student where name = 'BBB';
	err := GlobalConn.Unscoped().Where("name = ?", "陈8").Delete(new(Student)).Error
	fmt.Println(err)
}
