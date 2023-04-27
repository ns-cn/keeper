package main

import (
	"github.com/ns-cn/goter"
	"github.com/ns-cn/keeper/env"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

type CronCommand struct {
	Cron     string   `json:"cron"`
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
}

func main() {
	root := goter.NewRootCmdWithAction("keeper", "A simple tool like crontab", env.VERSION, func(command *cobra.Command, strings []string) {
		checkShell()
		pool := cron.New()
		updateInFiles(pool)
		go watchForUpdate(pool)
		pool.Start()
		select {}
	})
	root.Bind(&env.CfgFile, &env.CfgShell)
	_ = root.Execute()
}
