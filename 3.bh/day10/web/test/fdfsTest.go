package main

import (
	"fmt"
	"path"
)

func main()  {

	ret := path.Ext("hello.jpg")
	fmt.Println("ret=", ret[1:])
/*
	// 根据配置文件,初始化客户端
	clt, err := fdfs_client.NewClientWithConfig("/etc/fdfs/client.conf")
	if err != nil {
		fmt.Println("客户端初始化失败:", err)
		return
	}

	// 按文件名上传
	ret, err := clt.UploadByFilename("头像3.jpg")*/

	
}
