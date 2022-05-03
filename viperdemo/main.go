package main

// 使用viper 读取配置文件

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	// 读取yaml配置文件 并将配置文件映射到单层struct
	readYaml()

	// 读取yaml 多层配置信息并映射到嵌套结构体
	readMoreYaml()

	// 将线上线下配置文件进行隔离
	chooseConfig()

}

// 读取yaml配置文件 并将配置文件映射到单层struct
type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}

// 读取yaml 文件配置
func readYaml() {
	v := viper.New()

	// 指定文件路径
	v.SetConfigFile(".\\config\\config.yaml")

	// 查找阅读配置文件并处理错误信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置文件映射到struct
	sc := ServerConfig{}
	if err := v.Unmarshal(&sc); err != nil {
		panic(err)
	}
	fmt.Println(sc)

	// 直接输出配置文件信息
	name := v.Get("name")
	port := v.Get("port")
	fmt.Println(name)
	fmt.Println(port)
}

// 读取yaml 多层配置信息并映射到嵌套结构体

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type allServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

func readMoreYaml() {
	v := viper.New()

	// 指定文件路径
	v.SetConfigFile(".\\config\\config.yaml")

	// 查找阅读配置文件并处理错误信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)

	}

	// 将配置文件映射到struct
	allsc := allServerConfig{}
	if err := v.Unmarshal(&allsc); err != nil {
		panic(err)
	}
	fmt.Println(allsc)

}

// 将线上线下配置文件进行隔离
// 不用改任何代码 而且线上线下的配置文件能隔离开
// 通过环境变量实现此功能

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func chooseConfig() {
	debug := GetEnvInfo("CONFIG_DEBUG")
	configFileName := ".\\config\\config.yaml"

	if !debug {
		configFileName = ".\\config\\config_local.yaml"
	}

	v := viper.New()

	// 指定文件路径
	v.SetConfigFile(configFileName)

	// 查找阅读配置文件并处理错误信息
	if err := v.ReadInConfig(); err != nil {
		panic(err)

	}

	// 将配置文件映射到struct
	allsc := allServerConfig{}
	if err := v.Unmarshal(&allsc); err != nil {
		panic(err)
	}
	fmt.Println(allsc)

	// viper - 动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config uploadfile changed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&allsc)

		fmt.Println(allsc)

	})

	time.Sleep(time.Second * 200)
}
