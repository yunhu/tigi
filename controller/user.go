package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"tigi/client"
	"tigi/model"
	"tigi/st"
	"tigi/worker"
	"time"
)

//
func SyncToRedis(num interface{})  error{
	redisClient := client.GetRedisClient("goods")
	time.Sleep(time.Second * 3)
	set := redisClient.Set("taskId", num, time.Duration(0))
	return set.Err()
}

func GetUser(g *gin.Context) {
	id, _ := g.GetPostForm("id")
	cid, _:= strconv.Atoi(id)
	f:=SyncToRedis
		t  := &worker.Task{
			TaskId: cid,
			F:      f,
		}


	worker.Se


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
