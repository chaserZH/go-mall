package main

import (
	"fmt"
	"go-mall/conf"
)

func main() {

	loading() //加载配置
	fmt.Printf("hello world\n")

}

func loading() {
	conf.InitConfig()
	// 2. 打印一些配置信息验证配置加载成功
	fmt.Printf("当前运行环境: %s\n", conf.Config.System.AppEnv)
	fmt.Printf("服务监听端口: %s\n", conf.Config.System.HttpPort)
	fmt.Printf("MySQL主机: %s\n", conf.Config.Mysql["default"].DbHost)
}
