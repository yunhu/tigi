package model

import (
	"tigi/client"
	"tigi/st"
	"time"
)

func SyncToRedis(num interface{}) error {
	redisClient := client.GetRedisClient("goods")
	time.Sleep(time.Second * 10)
	set := redisClient.Set("taskId", num, time.Duration(0))
	return set.Err()
}
func AddUser(ust st.UserModel) int64 {

	mysqlClient := client.GetMysqlClient("frametest")
	create := mysqlClient.Create(ust)
	if create.Error != nil {
		return 0
	}
	return create.RowsAffected

	//创建
	//u := User{Name: "zyh", Age: 333, Sex: 1}
	//model.Mysql.Create(&u)

	//model.Mysql.LogMode(true)
	//
	////找单条
	//uu := User{Name: "zyh"}
	//var rr User
	//model.Mysql.First(&rr, &uu)
	//fmt.Println(rr)
	//rr.Name = "tiger"
	//model.Mysql.Save(&rr)
	//model.Mysql.Delete(&rr)
	//db.Model(&user).Where("active = ?", true).Update("name", "hello")
	//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	//找多条
	//var rr []User
	//model.Mysql.Where(map[string]interface{}{"name": "zyh", "sex": 1}).Find(&rr)
	//
	//fmt.Println(rr, uu)

}
