package controller

import (
	"github.com/gin-gonic/gin"
	"bj40ihome/web/utils"
	"net/http"
	"fmt"
	"image/png"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"

	getCaptcha "bj40ihome/web/proto/getCaptcha" // getCaptcha 为别名
	"context"
	"encoding/json"
	"github.com/afocus/captcha"

	userMicro "bj40ihome/web/proto/user" // userMicro 是别名
	"bj40ihome/web/model"
	"github.com/gomodule/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/tedcy/fdfs_client"
	"path"
)

// 获取 Session
func GetSession(ctx *gin.Context) {

	// 初始化容器,存储返回的错误信息
	//resp := make(map[string]string)
	resp := make(map[string]interface{})

	// 获取 Session数据
	s := sessions.Default(ctx)   // 初始化session 对象
	userName := s.Get("userName")

	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		var nameData struct{
			Name string `json:"name"`
		}

		nameData.Name = userName.(string)

		resp["data"] = nameData
	}

	ctx.JSON(http.StatusOK, resp)
}

// 获取 图片验证码
func GetImageCd(ctx *gin.Context) {
	// 获取 图片验证码uuid
	uuid := ctx.Param("uuid")

	// 指定服务发现
	consulReg := consul.NewRegistry()

	service := micro.NewService(micro.Registry(consulReg))

	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("go.micro.srv.getCaptcha", service.Client())

	// 调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Println("远程调用失败...")
		return
	}

	// 反序列化, 得到图片数据, 定义 开源 captcha 包下的 Image 变量.
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 生成 图片,显示回给浏览器
	png.Encode(ctx.Writer, img)

	fmt.Println("uuid = ", uuid)
}

// 获取 短信验证吗
func GetSmscd(ctx *gin.Context) {
	phone := ctx.Param("phone")
	imgCode := ctx.Query("text") // 获取用户动态输入的 验证码的数据值.
	uuid := ctx.Query("id")      // 获取 图片验证码uuid

	// 指定 consul 服务发现
	// 初始化 consul配置
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Registry(consulReg), // 指定使用的 服务发现
	)

	// 调用 Newxxx 初始化对象
	microClient := userMicro.NewUserService("go.micro.srv.user", service.Client())

	// 远程函数调用
	resp, _ := microClient.SendSms(context.TODO(), &userMicro.Request{Phone: phone, ImgCode: imgCode, Uuid: uuid})

	// 将微服务返回的数据, 发送给浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 注册用户
func PostRet(ctx *gin.Context) {
	// 创建 对象,用来存储 Request Payload 中的数据
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}

	ctx.Bind(&regData)

	// 创建service 对象
	service := utils.InitMicro()

	// 初始化
	microClient := userMicro.NewUserService("go.micro.srv.user", service.Client())

	// 调用远程函数
	resp, err := microClient.Register(context.TODO(), &userMicro.RegReq{
		Mobile:   regData.Mobile,
		Password: regData.PassWord,
		SmsCode:  regData.SmsCode,
	})

	if err != nil {
		fmt.Println("注册用户远程调用微服务失败...", err)
	}

	// 写回给前端
	ctx.JSON(http.StatusOK, resp)
}

// 获取地域信息
func GetArea(ctx *gin.Context) {

	// 从MySQL中获取数据,
	var areas []model.Area

	// 定义容器,写回数据给前端
	resp := make(map[string]interface{})

	// 获取redis连接池句柄, 得一条链接
	conn := model.RedisPool.Get()

	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))
	if len(areaData) == 0 { // 没有查询到 redis 中的数据.
		// 从MySQL中获取数据
		model.GlobalConn.Find(&areas)
		// 将数据序列化后,转换为 json 串, 写入到 redis 中.
		areaBuf, _ := json.Marshal(areas)
		// 把数据存储到 redis 中
		conn.Do("set", "areaData", areaBuf)
		fmt.Println("从 MySQL 中获取数....")
	} else {
		// 有数据, 做json 反序列化
		json.Unmarshal(areaData, &areas)
		fmt.Println("从 redis 中获取数....")
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	ctx.JSON(http.StatusOK, resp)
}

// 实现用户登录
func PostLogin(ctx *gin.Context) {

	// 封装容器, 存储写回给前端的数据.
	resp := make(map[string]interface{})

	// 获取 Request Payload 数据
	var loginData struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}
	// 利用 bind方法, 提取数据.
	ctx.Bind(&loginData)

	// 查询MySQL数据库, 校验用户输入的 mobile 和 password 是否正确.
	userName, err := model.Login(loginData.Mobile, loginData.Password)

	if err == nil {
		// 登录成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		// 将登录状态 保存到 session 中.
		// 获取一个 Session对象
		s := sessions.Default(ctx)
		// 设置Session
		s.Set("userName", userName)
		// 修改session内容时,必须要指定 save,才能生效.否则不生效.
		s.Save()

	} else {
		// 登录失败
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	}

	// 写回给浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 退出登录
func DeleteSession(ctx *gin.Context)  {
	// 初始化Session 对象
	s := sessions.Default(ctx)

	// 删除 session 数据, 该函数调用完成,并不生效
	s.Delete("userName")

	// 封装容器, 存储写回给前端的数据.
	resp := make(map[string]interface{})

	// 使用 save 保存
	err := s.Save()
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}

	// 写回给浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 获取用户信息
func GetUserInfo(ctx *gin.Context)  {
	// 获取当前用户 --- 获取session , 得"当前" 用户信息
	resp := make(map[string]interface{})

	defer ctx.JSON(http.StatusOK, resp)

	// 初始化Session 对象
	s := sessions.Default(ctx)
	userName := s.Get("userName")

	if userName == nil {   // 说明用户没登录, 异常. 判断为恶意进入.
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	} else {
		// 根据用户名 , 获取用户信息 --- 访问 mysql
		user, err := model.GetUserInfo(userName.(string))
		if err != nil {
			resp["errno"] = utils.RECODE_DBERR
			resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
			return
		}

		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		// 添加 map ,存储 剩余信息
		temp := make(map[string]interface{})
		temp["user_id"] = user.ID
		temp["name"] = user.Name
		temp["mobile"] = user.Mobile
		temp["real_name"] = user.Real_name
		temp["id_card"] = user.Id_card
		//temp["avatar_url"] = user.Avatar_url
		temp["avatar_url"] = "http://192.168.6.108:8888/" + user.Avatar_url

		resp["data"] = temp
	}
}

/*
get--查--获取数据
post--增--插入数据
put--改--更新数据
delete--删--删除数据
*/
// 更新用户名
func PutUserInfo(ctx *gin.Context)  {

	resp := make(map[string]interface{})

	defer ctx.JSON(http.StatusOK, resp)

	// 获取当前的用户名 -- 从当前的session 中获取.
	s := sessions.Default(ctx)
	userName := s.Get("userName")

	// 获取新用户名 --- 获取 Request Payload
	var nameData struct{
		Name string `json:"name"`
	}
	ctx.Bind(&nameData)

	// 更新用户名 --- update MySQL
	err := model.UpdateUserName(nameData.Name, userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	// 更新数据库成功后,必须要更新session 才能将更新实时的反应到 浏览器
	s.Set("userName", nameData.Name)
	s.Save()

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = nameData
}

// 上传头像
func PostAvatar(ctx *gin.Context)  {
	// 获取图片文件
	file, _ := ctx.FormFile("avatar")

	// 借助开源库,将文件存储到 fastDFS 中
	clt, _ := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")

	// 使用 file调用 Open函数, 打开文件读取文件内容
	f, _ := file.Open()

	// 创建缓冲区存储文件内容
	buf := make([]byte, file.Size)

	f.Read(buf)   // 不用 & 传参

	// 根据文件名,提取后缀名
	fileExt := path.Ext(file.Filename)

	// 上传图片文件字节流到 storage 中, 得到存储凭证
	remoteId, _ := clt.UploadByBuffer(buf, fileExt[1:])   // 参2 代表截掉 后缀中的 "."

	// 获取当前用户, 得到用户名
	s := sessions.Default(ctx)
	userName := s.Get("userName")   // 不需要校验, main.go 中 校验完成.

	// 根据用户名, 更新 MySQL中的 用户头像. 存入头像对应的凭证
	model.UpdateAvatar(userName.(string), remoteId)

	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["avatar_url"] = "http://192.168.6.108:8888/" + remoteId

	resp["data"] = temp
}

func PostUserAuth(ctx *gin.Context)  {
	//获取绑定数据
	var auth  struct {
		IdCard   string `json:"id_card"`
		RealName string `json:"real_name"`
	}
	err := ctx.Bind(&auth)
	//校验数据
	if err != nil {
		fmt.Println("获取数据错误", err)
		return
	}

	session := sessions.Default(ctx)
	userName := session.Get("userName")

	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	err = model.SaveRealName(userName.(string), auth.RealName, auth.IdCard)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
}

// 获取已发布的房屋信息 -- 伪实现
func GetUserHouses(ctx *gin.Context)  {
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = ""

	ctx.JSON(http.StatusOK, resp)
}

// 发布房源
func PostHouses(ctx *gin.Context)  {
	var house model.HouseInfo

	// gin 框架中的 bind获取数据时,不会转换数据类型. 前端产生的数据为 string,
	// 用来接受的 struct 成员也必须是 string 类型. 如果 int 报错!!!
	err := ctx.Bind(&house)
	if err != nil {
		fmt.Println("获取数据失败:", err)
		return
	}
	// 获取当前用户
	userName := sessions.Default(ctx).Get("userName")

	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)

	// 将当前的房屋信息, 插入到数据库中
	houseId, err := model.AddHouse(userName.(string), house)
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}

	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["house_id"] = houseId

	resp["data"] = temp
}
