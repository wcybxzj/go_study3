package main

import (
	"github.com/gin-gonic/gin"
	"bj40ihome/web/controller"
	"bj40ihome/web/model"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-contrib/sessions"
)

// 过滤用户登录状态 --- 校验 "当前" 用户
func LoginFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化 Session 对象
		s := sessions.Default(ctx)
		// 按 key 获取用户名
		userName := s.Get("userName")
		if userName == nil {		// 用户尚未登录
			ctx.Abort()  // 从这直接返回,不继续向后执行
		} else {
			ctx.Next()  // 继续下一个 中间件
		}
	}
}


func main()  {
	// 初始化 MySQL 连接池
	model.InitDb()

	// 初始化 redis 连接池
	model.InitRedis()

	// 初始化路由
	router := gin.Default()

	// 初始化容器 -- 给session
	store, _ := redis.NewStore(10, "tcp", "192.168.6.108:6379", "", []byte("bj40"))
	
	// 使用容器
	router.Use(sessions.Sessions("mysession", store))

	// 路由匹配
	router.Static("/home", "view")

	// 添加 路由分组
	r1 := router.Group("/api/v1.0/")  // v1.0 ---> v2.0
	{
		r1.GET("session", controller.GetSession)
		r1.GET("imagecode/:uuid", controller.GetImageCd)
		r1.GET("smscode/:phone", controller.GetSmscd)
		r1.POST("users/", controller.PostRet)
		r1.GET("areas", controller.GetArea)
		r1.POST("sessions", controller.PostLogin)

		r1.Use(LoginFilter())  // 从这开始, 以后的 Session 校验都不需要执行了.

		r1.DELETE("session", controller.DeleteSession)
		r1.GET("user", controller.GetUserInfo)
		r1.PUT("user/name", controller.PutUserInfo)
		r1.POST("user/avatar", controller.PostAvatar)

		r1.POST("user/auth",controller.PostUserAuth)
		r1.GET("user/auth",controller.GetUserInfo)

		// 获取以发布的房源
		r1.GET("user/houses", controller.GetUserHouses)

		// 发布房源
		r1.POST("houses", controller.PostHouses)
	}

	// 运行
	router.Run(":8087")
}
