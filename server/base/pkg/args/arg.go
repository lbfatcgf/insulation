package args

import "flag"

var ConfigPath, ConfigName *string
var Initialize *bool

func init() {
	ConfigPath = flag.String("confPath", "./config", "配置文件路径")
	ConfigName = flag.String("confName", "config.toml", "配置文件名称")
	Initialize = flag.Bool("init", false, "初始化,创建数据库表，超级管理员")
	flag.Parse()
}
