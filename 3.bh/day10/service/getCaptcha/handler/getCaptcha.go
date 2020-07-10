package handler

import (
	"context"
	getCaptcha "bj40ihome/service/getCaptcha/proto/getCaptcha"
	"github.com/afocus/captcha"
	"image/color"
	"encoding/json"
	"bj40ihome/service/getCaptcha/model"
	"fmt"
)

type GetCaptcha struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error {
	// 生成图片验证码

	// 1. 初始化对像
	cap := captcha.New()

	// 2. 设置字体  -- 需要导入字体库
	cap.SetFont("./conf/comic.ttf")

	// 3. 设置验证区域大小
	cap.SetSize(128, 64)
	// 4. 设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)
	// 5. 设置前景色
	cap.SetFrontColor(color.RGBA{255,255,255,255})
	// 6. 设置背景色
	cap.SetBkgColor(color.RGBA{10,255,0,128}, color.RGBA{255,128,64,32})

	// 生成图片通验证码
	img, str := cap.Create(4, captcha.NUM)

	fmt.Println("--------------str = ", str)
	// 调用函数写入数据库
	err := model.SaveImgCode(str, req.Uuid)
	if err != nil {
		fmt.Println("微服务存储redis 失败:", err)
		return err
	}

	// 将生成的图片,序列化
	imgBuf, _ := json.Marshal(img)

	// 借助 Response 将 序列化后的 图片 传出
	rsp.Img = imgBuf

	return nil
}
