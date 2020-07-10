package main

import (
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main()  {
	// 1. 初始化对像
	cap := captcha.New()

	// 2. 设置字体  -- 需要导入字体库
	cap.SetFont("./comic.ttf")

	// 3. 设置验证区域大小
	cap.SetSize(128, 64)
	// 4. 设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)
	// 5. 设置前景色
	cap.SetFrontColor(color.RGBA{255,255,255,255})
	// 6. 设置背景色
	cap.SetBkgColor(color.RGBA{10,255,0,128}, color.RGBA{255,128,64,32})
	// 7. 展示图片验证码到页面
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cap.Create(4, captcha.NUM)
		png.Encode(w, img)
		println(str)
	})

	// 启动服务
	http.ListenAndServe(":8123", nil)
}
