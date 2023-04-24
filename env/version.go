package env

import "github.com/ns-cn/goter"

const VERSION = "1.00.2023.0424"

var (
	CfgFile = goter.NewCmdFlagString("", "load", "l", "config file (default is keeper.json)")
)
