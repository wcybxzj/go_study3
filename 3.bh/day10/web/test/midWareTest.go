package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)

// 创建 中间件 1
func Test1(ctx *gin.Context)  {

	fmt.Println("111")

	t := time.Now()   // 获取时间

	ctx.Next()

	fmt.Println(time.Now().Sub(t))

	fmt.Println("444")
}

// 创建 中间件 2
func Test2() gin.HandlerFunc {
	return func(context *gin.Context) {

		fmt.Println("333")

		//context.Abort()
		context.Next()

		fmt.Println("555")
	}
}

func main() {
	router := gin.Default()

	// 使用中间件1
	router.Use(Test1)  // 不传 Test1()

	// 使用中间件2
	router.Use(Test2())  // 传参方式不同, 该格式需要添加 ().

	router.GET("/test", func(context *gin.Context) {
		fmt.Println("222")
		context.Writer.WriteString("hello middleWare")
	})

	router.Run(":9876")
}
