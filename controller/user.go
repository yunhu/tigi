package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"tigi/common"
)

//
func GetUser(g *gin.Context) {
	g.JSON(200, "")
}
func AddUser(g *gin.Context) {
	u := common.User{}
	ID, _ := g.GetPostForm("ID")
	Name, _ := g.GetPostForm("name")
	Age, _ := g.GetPostForm("age")
	Sex, _ := g.GetPostForm("sex")

	id, _ := strconv.Atoi(ID)
	u.ID = id
	u.Name = Name
	age, _ := strconv.Atoi(Age)
	u.Age = age
	sex, _ := strconv.Atoi(Sex)
	u.Sex = sex
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code": 33,
			"msg":  err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": u,
	})
}
