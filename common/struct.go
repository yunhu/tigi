package common

//定义结构

//用户结构
type User struct {
	ID   int    `form:"ID" validate:"required,min=5,max=10"`
	Name string `form:"name" validate:"required,min=3,max=11"`
	Age  int    `form:"age" validate:"gte=1,lte=100"`
	Sex  int    `form:"sex" validate:"gte=0,lte=1"`
}

func (u User) TableName() string {
	return "user"
}
