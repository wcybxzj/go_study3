package main

import (
	"os/exec"
	"fmt"
)

//例子1:最简单的golang执行cmd 不获取命令输出
func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	// cmd = exec.Command("/bin/bash", "-c", "echo 1;echo2;")

	cmd = exec.Command("/bin/bash", "-c", "echo 123 > /tmp/123.txt")

	err = cmd.Run()

	fmt.Println(err)
}
