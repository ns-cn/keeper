package env

import "github.com/ns-cn/goter"

var (
	CfgFile  = goter.NewCmdFlagString("keeper.json", "load", "l", "config file")
	CfgShell = goter.NewCmdFlagString("bash", "shell", "s", "running with specific shell")
)
