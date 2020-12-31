package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tigi/client"
	"tigi/common"
	"tigi/router"
)

func main() {
	//初始化配置
	config := common.GetConfig()
	//初始化client
	client.InitClient(&config)
	//gin 初始化引擎
	engine := gin.Default()
	//注册路由
	router.Register(engine)
	ports := strconv.Itoa(config.Port)
	//启动服务
	engine.Run(":" + ports)

}
