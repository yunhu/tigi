package router
import (
	"tigi/controller"
	"net/http"
)

func  Register()  {
	ctl := controller.Ctl{}
	http.HandleFunc("/index", ctl.Index)
	http.HandleFunc("/", ctl.Index)
}
