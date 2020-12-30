package model

type User struct {
	ID   int
	Name string
	Age  int
	Sex  int
}

func (u User) TableName() string {
	return "user"

}

func Test() {


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
