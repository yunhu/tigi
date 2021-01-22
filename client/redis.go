package client

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"tigi/config"
)

const (
	CheckRedisNum = 1
)

var RedisHandle map[string]*redis.Client = nil

//初始化redis
func InitRedis(conf *config.Conf) {
	if RedisHandle == nil {
		OpenRedis(conf)
	}
}

//获取client redis
func GetRedisClient(DbName string) *redis.Client {
	if RedisHandle[DbName] == nil {
		panic(DbName + "未初始化！")
	}
	return RedisHandle[DbName]
}

//检查配置
func CheckRedisConf(conf *config.Conf) bool {
	if len(conf.Redis[0].IP) > CheckRedisNum {
		return true
	}
	return false
}

func OpenRedis(conf *config.Conf) {
	if len(conf.Redis[0].IP) < 3 {
		return
	}
	for _, v := range conf.Redis {
		rc := redis.NewClient(&redis.Options{
			Addr:     v.IP,
			Password: v.Pass,
			DB:       int(v.DbIndex),
		})
		err := rc.Ping().Err()

		if err != nil {
			logrus.Error("Redis init error:", err.Error())
			//			panic("redis 初始化失败")
		} else {
			RedisHandle = make(map[string]*redis.Client)
			RedisHandle[v.DbName] = rc
		}
	}
}
