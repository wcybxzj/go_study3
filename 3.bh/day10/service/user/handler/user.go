package handler

import (
	"context"
	user "bj40ihome/service/user/proto/user"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"bj40ihome/service/user/model"
	"math/rand"
	"time"
	"fmt"
	"bj40ihome/service/user/utils"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) SendSms(ctx context.Context, req *user.Request, rsp *user.Response) error {
	// 校验 图片验证码是否正确
	result := model.CheckImgCode(req.ImgCode, req.Uuid)
	if result {
		// 校验成功, 发送短信验证码
		client, _ := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4G8Y5UpJvFihufkcHw5k", "cve7fFn95vAGIRsxxgP2dDu8s4sy9g")

		request := dysmsapi.CreateSendSmsRequest()
		request.Scheme = "https"
		request.Domain = "dysmsapi.aliyuncs.com"  // 酌情添加.

		// 播种随机数种子
		rand.Seed(time.Now().UnixNano())
		// 生成随机数.作短信验证码
		smsCode := fmt.Sprintf("%06d", rand.Int31n(1000000))

		request.PhoneNumbers = req.Phone
		request.SignName = "爱家租房网"
		request.TemplateCode = "SMS_183242785"
		request.TemplateParam = `{"code":"` + smsCode + `"}`

		// 发送短信验证码
		response, _ := client.SendSms(request)
		if response.IsSuccess() {

			// 将短信验证码,写入到 redis 中, 以 phone_code 作为key
			err := model.SaveSmsCode(req.Phone, smsCode)
			if err != nil {
				fmt.Println("存储短信验证码到 redis 失败:", err)
				rsp.Errno = utils.RECODE_SMSERR
				rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
			}
			// 显示成功信息
			fmt.Println("生成的短信验证码为:", smsCode)
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		} else {
			rsp.Errno = utils.RECODE_SMSERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_SMSERR)
		}
		fmt.Printf("response is %#v\n", response)

	} else {
		// 校验失败
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}

	return nil
}

// 实现注册用户
func (e *User) Register(ctx context.Context, req *user.RegReq, rsp *user.Response) error {

	fmt.Println("-------------------0")

	// 校验 redis 中存储的短信验证码, 是否正确.
	err := model.CheckSmsCode(req.Mobile, req.SmsCode)

	fmt.Println("-------------------1")

	if err == nil {
		fmt.Println("-------------------2")

		// 如果正确,完成用户注册, 将数据写入到MySQL对应表. --- user表
		err = model.RegisterUser(req.Mobile, req.Password)
		if err != nil {
			fmt.Println("-------------------3")
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg= utils.RecodeText(utils.RECODE_DBERR)
		} else {
			fmt.Println("-------------------4")
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg= utils.RecodeText(utils.RECODE_OK)
		}

	} else {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg= utils.RecodeText(utils.RECODE_DATAERR)
	}
	fmt.Println("-------------------5")

	return nil
}