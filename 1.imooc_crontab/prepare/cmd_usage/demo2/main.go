package main

import (
	"os/exec"
	"fmt"
)

//例子2:golang执行cmd 获取命令输出
func main() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	// 生成Cmd
	cmd = exec.Command(
		"/bin/bash",
		"-c",
		"/usr/bin/php /root/www/go_www/src/go_study2/2.imooc_crontab/prepare/cmd_usage/demo2/1.php")

	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println("error:"+err.Error())
		return
	}

	// 打印子进程的输出
	fmt.Println(string(output))
}