package client

import "tigi/config"

//初始化,各个外部连接 等mysql redis
func InitClient(conf *config.Conf) {
	//mysql 初始化
	if CheckMysqlConf(conf) {
		InitMysql(conf)
	}

	//redis 初始化
	if CheckRedisConf(conf) {
		InitRedis(conf)
	}
}
