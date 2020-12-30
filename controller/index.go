package controller

import (
	"github.com/gin-gonic/gin"
)

func Index(g *gin.Context) {
	g.String(200, "ok")
	//common.Msg(w,0,"hello","")
}
