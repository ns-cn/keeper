package main

import (
	"encoding/json"
	"fmt"
	"github.com/ns-cn/keeper/env"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func updateInFiles(cron *cron.Cron) {
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
	for _, cronCommand := range CronCommands {
		var name = cronCommand.Name
		commands := cronCommand.Commands
		_, err := cron.AddFunc(cronCommand.Cron, func() {
			fmt.Printf("%v : %s\n", time.Now(), name)
			for _, cmd := range commands {
				execCommand(cmd)
			}
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
