package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func conf() (string, string) {
	viper.SetConfigName("config") //设置配置文件的名字
	viper.AddConfigPath(".")      //添加配置文件所在的路径
	viper.SetConfigType("yaml")   //设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	channelS := viper.GetString("channel.secret")
	channelT := viper.GetString("channel.token")
	return channelS, channelT
}

func main() {
	channelS, channelT := conf()
}
