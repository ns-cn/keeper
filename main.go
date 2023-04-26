package main

import (
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/ns-cn/keeper/env"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"os/exec"
)

// 根据一个字符串执行命令
func execCommand(command string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	cmd := exec.Command("bash", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type CronCommand struct {
	Cron     string   `json:"cron"`
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}

func main() {
	root := goter.NewRootCmdWithAction("keeper", "A simple tools like crontab", env.VERSION, func(command *cobra.Command, strings []string) {
		if env.CfgFile.Value == "" {
			env.CfgFile.Value = "keeper.json"
		}
		cron := cron.New()
		updateInFiles(cron)
		go watchForUpdate(cron)
		cron.Start()
		select {}
	})
	root.Bind(&env.CfgFile)
	_ = root.Execute()
}
