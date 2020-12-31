package common

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"tigi/config"
)

const (
	go_develop_mode = "GODEVELOPMODE"
)

//sting转md5
func Md5(str string) string {
	m := crypto.MD5.New()
	io.WriteString(m, str)
	md5 := fmt.Sprintf("%x", m.Sum(nil))
	return md5

}

//byte类型 转md5
func Md5Byte(b []byte) string {
	m := crypto.MD5.New()
	m.Write([]byte(b))
	return hex.EncodeToString(m.Sum(nil))
}

//获取配置
func GetConfig() config.Conf {
	conf := config.Conf{}
	dir, _ := os.Getwd()
	mode := os.Getenv(go_develop_mode)
	if mode != "dev" {
		mode = "online"
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("现在默认是线上环境，如果要切换到线下请执行 export GODEVELOPMODE=dev")
	} else {
		gin.SetMode(gin.DebugMode)
		fmt.Println("现在是在测试环境")
	}
	filePath := dir + "/config/config.toml." + mode
	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		panic(err)
	} else {
		return conf
	}

}
