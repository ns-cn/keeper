package main

import (
	"encoding/json"
	"fmt"
	"github.com/ns-cn/goter"
	"github.com/ns-cn/keeper/env"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"time"
)

// 根据一个字符串执行命令
func execCommand(cron CronCommand, command string) {
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
		var CronCommands = make([]CronCommand, 0)
		data, err := os.ReadFile(env.CfgFile.Value)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &CronCommands)
		if err != nil {
			panic(err)
		}
		fmt.Println(CronCommands)
		// 每小时执行一次
		cronPool := cron.New()
		for _, cronCommand := range CronCommands {
			_, err := cronPool.AddFunc(cronCommand.Cron, func() {
				fmt.Printf("%v : %s\n", time.Now(), cronCommand.Name)
				for _, cmd := range cronCommand.Commands {
					execCommand(cronCommand, cmd)
				}
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		cronPool.Start()
		select {}
	})
	root.Bind(&env.CfgFile)
	_ = root.Execute()
}
