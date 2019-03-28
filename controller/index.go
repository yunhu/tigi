package controller

import (
	"fmt"
	"net/http"
)

func (this *Ctl) Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"hello")
	//common.Msg(w,0,"hello","")
}
