package args

import "flag"

var ConfigPath, ConfigName *string

func init() {
	ConfigPath = flag.String("confPath", "./config", "config path")
	ConfigName = flag.String("confName", "config.toml", "config name")
	flag.Parse()
}
