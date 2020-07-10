package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
	"fmt"
)

func main()  {
	router := gin.Default()

	// 初始化容器
	store, _ := redis.NewStore(10, "tcp", "192.168.6.108:6379", "", []byte("bj40"))

	// 设置临时Session
	//store.Options(sessions.Options{MaxAge:0})

	// 使用容器
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/test", func(context *gin.Context) {
		context.SetCookie("testCookie", "chuanzhi-itcast", 0, "", "", false, true)

		cookieVal, _ := context.Cookie("testCookie")

		fmt.Println("获取到:", cookieVal)
		// 获取一个 Session对象
		s := sessions.Default(context)

		// 设置Session
		s.Set("itcast", "itheima123")

		// 修改session内容时,必须要指定 save,才能生效.否则不生效.
		s.Save()

		// 获取session
		v := s.Get("itcast")

		// 将返回的 Interface 类型的 v, 断言成string 类型, 输出.
		fmt.Println("获取的session:", v.(string))

		context.Writer.WriteString("Session测试开始了....")
	})

	router.Run(":9876")
}
