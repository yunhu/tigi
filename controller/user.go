package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"tigi/model"
	"tigi/st"
	"tigi/worker"
)

//


//测试协程池的使用
func GetUser(g *gin.Context) {
	id, _ := g.GetPostForm("id")
	cid, _:= strconv.Atoi(id)
	f:=model.SyncToRedis
		t  := &worker.Task{
			TaskId: cid,
			F:      f,
		}


	worker.WP.SengToWorker(t)

	g.JSON(200, id)
}
func AddMem(g *gin.Context) {
	u := st.UserModel{}
	Name, _ := g.GetPostForm("name")
	Age, _ := g.GetPostForm("age")
	Sex, _ := g.GetPostForm("sex")
	u.Name = Name
	age, _ := strconv.ParseInt(Age, 10, 64)
	u.Age = uint64(age)
	sex, _ := strconv.ParseInt(Sex, 10, 64)
	u.Sex = uint64(sex)
	v := validator.New()
	err := v.Struct(u)
	//logrus.Infof("adduser方法", u)
	if err != nil {
		g.JSON(http.StatusOK, gin.H{
			"code": 33,
			"msg":  err.Error(),
		})
		return
	}
	if model.AddUser(u) > 0 {
		g.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "ok",
			"data": u,
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"code": 124,
			"msg":  err,
		})
	}

}
