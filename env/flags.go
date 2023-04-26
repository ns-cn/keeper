package env

import "github.com/ns-cn/goter"

var (
	CfgFile = goter.NewCmdFlagString("", "load", "l", "config file (default is keeper.json)")
)
