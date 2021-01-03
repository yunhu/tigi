package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"tigi/client"
	"tigi/common"
	"tigi/router"
	"time"
)

func main() {
	//关闭控制台颜色
	gin.DisableConsoleColor()
	//初始化配置
	config := common.GetConfig()
	//初始化client
	client.InitClient(&config)
	//gin 初始化引擎
	engine := gin.Default()
	//engine.Use(middleware.LogerMiddleware())
	//注册路由
	router.Register(engine)
	ports := strconv.Itoa(config.Port)
	//启动服务
	srv := http.Server{
		Addr:    ":" + ports,
		Handler: engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("httperror!")
		}
	}()
	//平滑重启
	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, os.Interrupt, syscall.SIGUSR2)
	<-quitSignal
	fmt.Println("开始关闭")
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("close error!")
	}
	fmt.Println("server closed!")

}
