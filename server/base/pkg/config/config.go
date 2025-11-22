package config

import (
	"fmt"

	"insulation/server/base/pkg/jwt_util"

	"github.com/spf13/viper"
)

var (
	filePath = "./config"
	fileName = "config"
)

var (
	globalConfig *GlobalConfig
	viperReader  *viper.Viper
)

type GlobalConfig struct {
	Web struct {
		Port int    `mapstructure:"port" toml:"port"`
		Name string `mapstructure:"name" toml:"name"`
	} `mapstructure:"web" toml:"web"`
	Mode string `mapstructure:"mode" toml:"mode"`
	Log  struct {
		Level string `mapstructure:"level" toml:"level"`
		Path  string `mapstructure:"path" toml:"path"`
	} `mapstructure:"log" toml:"log"`
	DataSource struct {
		DataBase struct {
			DSN string `mapstructure:"dsn" toml:"dsn"`
		} `mapstructure:"dataBase" toml:"dataBase"`
		Redis struct {
			DSN string `mapstructure:"dsn" toml:"dsn"`
		} `mapstructure:"redis" toml:"redis"`
	} `mapstructure:"dataSource" toml:"dataSource"`
	JwtOptions jwt_util.JwtOptions `mapstructure:"jwt_option" toml:"jwt_option"`
}

func CustomConfig() *viper.Viper {
	if viperReader == nil {
		panic("config has not initialize")
	}
	return viperReader
}

func Global() *GlobalConfig {
	if globalConfig == nil {
		panic("config has not initialize")
	}
	return globalConfig
}

func DefaultInitialize() {
	Initialize(filePath, fileName)
}

func Initialize(fPath, fName string) {
	v := viper.New()
	v.SetConfigName(fName)
	v.AddConfigPath(fPath)
	v.SetConfigType("toml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	globalConfig = &GlobalConfig{}
	err = v.Unmarshal(globalConfig)
	if err != nil {
		panic(err)
	}
	viperReader = v
}

func IsDebug() bool {
	fmt.Println(globalConfig.Mode == "debug")
	return globalConfig.Mode == "debug"
}

func IsRelease() bool {
	return globalConfig.Mode == "release"
}

func IsTest() bool {
	return globalConfig.Mode == "test"
}
