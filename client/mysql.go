package client

import (
	"fmt"
	"gorm.io/gorm"
	"tigi/config"
)

const (
	CheckMysqlNum int = 1
)

var MysqlHandle map[string]*gorm.DB = nil

//初始化mysql
func InitMysql(conf *config.Conf) {
	if MysqlHandle == nil {
		OpenMysql(conf)
	}
}

//获取mysql实例
func GetMysqlClient(DbName string) *gorm.DB {
	if MysqlHandle[DbName] == nil {
		panic(DbName + "未初始化！")
	}
	return MysqlHandle[DbName]
}

//单个开

//连接mysql
func OpenMysql(conf *config.Conf) {

	for _,v:= range conf.Mysql{
		fmt.Println(v)
	}
/*		timeout := strconv.FormatInt(v.Timeout, 10)

		dsn := v.User + ":" + v.Pass + "@(" + v.IP + ")/" + v.DbName + "?charset=utf8&parseTime=True&loc=Local&timeout="+timeout+"s"
		db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
		if err != nil {
			panic(err)
		}else{
			MysqlHandle[v.DbName] = db
		}
	}*/
}

//检测mysql配置
func CheckMysqlConf(conf *config.Conf) bool {
	if len(conf.Mysql[0].User) > CheckMysqlNum && len(conf.Mysql[0].IP) > CheckMysqlNum && len(conf.Mysql[0].Pass) > CheckMysqlNum {
		return true
	}
	return false
}
