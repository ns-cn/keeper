# keeper
一个简易的定时任务处理程序

## 介绍
- [x] 无需重启服务，修改配置文件即刻生效
- [x] 可指定命令参数运行的shell
- [x] 灵活配置多任务（一个定时任务多个任务，可配置多个定时任务）

## 配置文件

```json
[
  {
    "cron": "10 * * * *",
    "name": "测试周期执行1",
    "commands": [
      "ls -al",
      "echo hello"
    ]
  },
  {
    "cron": "20 * * * *",
    "name": "测试周期执行2",
    "commands": [
      "ls -al"
    ]
  }
]
```

## 使用

```bash
keeper --help
# 默认使用./keeper.json作为配置文件
keeper
# 指定引用的配置文件
keeper --load other_keeper.json
# 指定shell运行命令
keeper --shell /bin/bash --load other_keeper.json
```

```shell
A simple tool like crontab

Usage:
  keeper [flags]
  keeper [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     打印当前版本号

Flags:
  -h, --help           help for keeper
  -l, --load string    config file (default "keeper.json")
  -s, --shell string   running with specific shell (default "bash")

Use "keeper [command] --help" for more information about a command.
```