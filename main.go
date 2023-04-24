package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"time"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "keeper",
	Short: "A simple tools like crontab",
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			cfgFile = "keeper.json"
		}
		var CronCommands = make([]CronCommand, 0)
		data, err := os.ReadFile(cfgFile)
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
		for _, cron := range CronCommands {
			_, err := cronPool.AddFunc(cron.Cron, func() {
				fmt.Printf("%v : %s\n", time.Now(), cron.Name)
				for _, cmd := range cron.Commands {
					execCommand(cron, cmd)
				}
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		cronPool.Start()
		select {}
	},
}

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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "load", "l", "", "config file (default is keeper.json)")
	_ = rootCmd.Execute()
}
