package model

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tigi/common"
	"tigi/config"
)

type Model struct {
	Redis *redis.Client
	Mysql *gorm.DB
}

func mysql_init(conf config.Conf) *gorm.DB {
	dsn := conf.Mysql.User + ":" + conf.Mysql.Pass + "@(" + conf.Mysql.IP + ")/" + conf.Mysql.DB + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

func redis_init(conf config.Conf) *redis.Client {
	if len(conf.Redis.IP) < 3 {
		return nil
	}
	rc := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.IP,
		Password: conf.Redis.Pass,
		DB:       conf.Redis.DB,
	})
	return rc
}

func New() *Model {
	conf := common.GetConfig()
	return &Model{Redis: redis_init(conf), Mysql: mysql_init(conf)}
}
