package router

import (
	"github.com/gin-gonic/gin"
	"tigi/controller"
)

func Register(engine *gin.Engine) {
	//===================用户部分==================================
	engine.POST("/user/add", controller.AddUser)
	engine.POST("/user/list", controller.GetUser)
	//===================用户部分==================================

	//===================权限部分==================================
	//===================权限部分==================================

	//重定向
	//engine.GET("/redirect", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/index")
	//})

}
