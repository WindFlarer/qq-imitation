package main

import (
	"wechat-imitation/initialization"
	"wechat-imitation/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化
	initialization.InitAll()
	// 创建引擎
	r := gin.Default()
	//加载路由
	router.InitRouters(r)
	// 启动
	r.Run()
}
