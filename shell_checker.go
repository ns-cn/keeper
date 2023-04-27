package main

import (
	"fmt"
	"github.com/ns-cn/keeper/env"
	"os/exec"
)

func checkShell() {
	err := exec.Command(env.CfgShell.Value, "--help").Run()
	if err != nil {
		panic(fmt.Sprintf("%s is not a shell executor\n%v", env.CfgShell.Value, err))
	}
}

// 根据一个字符串执行命令
func execCommand(command string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	cmd := exec.Command(env.CfgShell.Value, "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
}
