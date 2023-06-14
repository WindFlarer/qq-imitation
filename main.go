package main

import (
	"fmt"
	"wechat-imitation/config"
	"wechat-imitation/initialization"
)

func main() {
	// 初始化
	initialization.LoadConfig() // 读取yaml
	initialization.InitDB()     // 连接数据库
	initialization.InitTable()  // 初始化表
	fmt.Println(config.Conf.Mysql.Username)
}
