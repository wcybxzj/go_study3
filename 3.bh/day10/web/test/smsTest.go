package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

)

func main() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4G8Y5UpJvFihufkcHw5k", "cve7fFn95vAGIRsxxgP2dDu8s4sy9g")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.Domain = "dysmsapi.aliyuncs.com"  // 酌情添加.

	request.PhoneNumbers = "18610382737"
	request.SignName = "爱家租房网"
	request.TemplateCode = "SMS_183242785"
	//request.TemplateParam = "{\"code\":\"123454321\"}"
	request.TemplateParam = `{"code":"123454321"}`

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

