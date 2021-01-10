package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"tigi/client"
	"tigi/common"
	"tigi/router"
	"tigi/worker"
	"time"
)

func main() {
	defer func() {
		if err := recover();err !=nil{
			log.Fatal("出错了:",err)
		}
	}()
	// 启动多核调度
	runtime.GOMAXPROCS(runtime.NumCPU())
	//关闭控制台颜色
	gin.DisableConsoleColor()
	//初始化配置
	config := common.GetConfig()
	//初始化client
	client.InitClient(&config)
	//gin 初始化引擎
	engine := gin.Default()
	//注册路由
	router.Register(engine)
	ports := strconv.Itoa(config.Port)
	//协程
	p:=worker.NewPool(5)
	p.Run()

	//启动服务
	srv := http.Server{
		Addr:    ":" + ports,
		Handler: engine,
	}
	go func() {
	    fmt.Println("before")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("httperror!")
		}
		fmt.Println("start " +ports)
	}()
	//平滑重启
	quitSignal := make(chan os.Signal)
	signal.Notify(quitSignal, os.Interrupt, syscall.SIGUSR2)
	fmt.Println("sleep")
	time.Sleep(time.Second * 5)
	<-quitSignal
	fmt.Println("开始关闭")
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("close error!")
	}
	fmt.Println("server closed!")

}
