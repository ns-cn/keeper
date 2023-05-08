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
func execCommand(command string, cron CronCommand) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var cmd *exec.Cmd
	if cron.SpecificShell {
		if cron.Shell == "" {
			cmd = exec.Command(command)
		} else {
			cmd = exec.Command(cron.Shell, "-c", command)
		}
	} else {
		cmd = exec.Command(env.CfgShell.Value, "-c", command)
	}
	if cron.Dir != "" {
		cmd.Dir = cron.Dir
	}
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
}
