package main

import (
	"fmt"
	"net/http"
	"strconv"
	"tigi/common"
	"tigi/router"
)

func main() {

	//初始化配置
	config := common.GetConfig()
	//注册路由
	router.Register()
	ports := strconv.Itoa(config.Server.Port)
	fmt.Println("start with port " + ports)
	//启动服务
	http.ListenAndServe(":"+ports, nil)
}
