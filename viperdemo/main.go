package main

// 使用viper 读取配置文件

import (
	"fmt"

	"github.com/spf13/viper"
)





func main(){
	readYaml()

}


// 将配置文件映射到struct
type ServerConfig struct{
	ServiceName string `mapstructure:"name"`
	Port int `mapstructure:"port"`
}




// 读取yaml 文件配置
func readYaml(){
	v := viper.New()

	// 指定文件路径
	v.SetConfigFile(".\\config\\config.yaml")


	// 查找阅读配置文件并处理错误信息
	if	err := v.ReadInConfig();  err != nil{
		panic(err)
	}


	// 将配置文件映射到struct
	sc := ServerConfig{}
	if err := v.Unmarshal(&sc); err != nil{
		panic(err)
	}
	fmt.Println(sc)

	// 直接输出配置文件信息
	name := v.Get("name")
	port := v.Get("port")
	fmt.Println(name)
	fmt.Println(port)
} 