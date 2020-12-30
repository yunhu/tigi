package common

import (
	"crypto"
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

func Md5(str string) string {

	m := crypto.MD5.New()
	io.WriteString(m, str)
	md5 := fmt.Sprintf("%x", m.Sum(nil))
	return md5

}

func GetConfig() config.Conf {
	conf := config.DB{}
	dir, _ := os.Getwd()
	filePath := dir + "/config/config.toml"
	mode := os.Getenv(go_develop_mode)
	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		panic(err)
	} else {
		if mode == "dev" {
			gin.SetMode(gin.DebugMode)
			fmt.Println("现在是在测试环境")

			return conf.Dev
		} else {
			gin.SetMode(gin.ReleaseMode)
			fmt.Println("现在默认是线上环境，如果要切换到线下请执行 export GODEVELOPMODE=dev")
			return conf.Online
		}
	}

}

