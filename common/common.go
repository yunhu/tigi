package common

import (
	"tigi/config"
	"crypto"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"net/http"
	"os"
)


const (
	go_develop_mode = "GODEVELOPMODE"
	config_path = "/src/4mq/config/config.toml"
	gopath = "GOPATH"

)
func Md5(str string) string {


	m := crypto.MD5.New()
	io.WriteString(m, str)
	md5 := fmt.Sprintf("%x", m.Sum(nil))
	return md5

}

func GetConfig()  config.Conf{
	conf := config.DB{}
	path := os.Getenv(gopath)
	filePath := path + config_path
	mode := os.Getenv(go_develop_mode)
	if _,err := toml.DecodeFile(filePath,&conf); err != nil {
		panic(err)
	}else{
		if mode == "dev" {
			return conf.Dev
		}else {
			return conf.Online
		}
	}

}

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Msg(w http.ResponseWriter, code int, msg string, arr interface{}) {
	data := &Message{
		Code: code,
		Msg:  msg,
		Data: arr,
	}
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}