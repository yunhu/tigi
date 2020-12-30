package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tigi/common"
	"tigi/router"
)

func main() {
	//初始化配置
	config := common.GetConfig()
	//gin 初始化引擎
	engine := gin.Default()
	//注册路由
	router.Register(engine)
	ports := strconv.Itoa(config.Server.Port)
	//启动服务
	engine.Run(":" + ports)

}
