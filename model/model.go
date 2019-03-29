package model

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	Redis *redis.Client
	Db    *gorm.Model
}

func New() *Model {
	db,err := gorm.Open("mysql","user:password@/dbname?charset=utf8&parseTime=True&loc=Local)
	if err != nil {
		panic(err)
	}
	defer func() {
		db.Close()
	}()

	return &Model{}
}
