package router

import (
	"github.com/gin-gonic/gin"
	"tigi/controller"
)


func  Register(engine *gin.Engine)  {
	engine.POST("/t",controller.Index)

	//重定向
	//engine.GET("/redirect", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/index")
	//})

}
